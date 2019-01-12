package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"log"
	//"strings"
	"errors"
	"encoding/json"
	//"github.com/fatih/color"
	//"bytes"
)

type Client struct {
	id      	int
	message 	string
	conn    	net.Conn
	username	string
	target 		string
}

type Request struct {
	Reqtag string `json:"reqtag"`
	Username string `json:"username"`
	Pubkey string `json:"pubkey"`
	Message string `json:"message"`
}

const PREFIX 			= "~&#"
const SUFFIX			= "#&~"
const GET_CLIENTS   	= PREFIX + "get_clients" + SUFFIX
const LOGIN 	    	= PREFIX + "login" + SUFFIX
const SIGNUP 	    	= PREFIX + "signup" + SUFFIX
const SELECT_TARGET 	= PREFIX + "selectTarget" + SUFFIX
const TARGET_FAIL   	= PREFIX + "targetFail" + SUFFIX
const TARGET_SET    	= PREFIX + "targetset" + SUFFIX
const SIGNUP_FAILURE    = PREFIX + "signupfailure" + SUFFIX
const SIGNUP_SUCCESSFUL = PREFIX + "signupsuccess" + SUFFIX
const ERROR    			= PREFIX + "error" + SUFFIX

var clientsList []Client
//var clientsList = make(map[string]Client)

func requestHandler(client Client){

	request := Request{}
	json.Unmarshal([]byte(client.message), &request)

	switch request.Reqtag {
	case GET_CLIENTS:
			clients := ""
			for _, client := range clientsList {
			    clients = clients+" : "+client.username
			}
			go fmt.Fprintf(client.conn, clients+"\n")

		break
	case LOGIN:
			go fmt.Fprintf(client.conn, "logged in "+"\n")

		break
	case SIGNUP:
			err := signup(client, request.Username)
			if err != nil{
				go fmt.Fprintf(client.conn, SIGNUP_FAILURE+"\n")
			}else{
				go fmt.Fprintf(client.conn, ERROR+" : "+SIGNUP_SUCCESSFUL+"\n")
			}

		break
	case SELECT_TARGET:
			err := setTarget(client, request.Username)
			if err != nil{
				go fmt.Fprintf(client.conn, TARGET_FAIL+"\n")
			}else{
				go fmt.Fprintf(client.conn, TARGET_SET+"\n")
				//plus attach the public key to the json
			}

		break
		default:
			go fmt.Fprintf(client.conn, "OKAY . . ."+"\n")

		break
	}
}

func setTarget(client Client, username string) error{
	//set this username as the target for this client
	return nil
}

func userExists(username string) bool{
	for _, client := range clientsList {
	    if username == client.username{
				return true
			}
	}
	return false
}

func signup(client Client, username string) error{
	//check for existing username and
	//send either ~&#signupsuccess#&~
	//or ~&#error#&~ + ~&#signupfailure#&~
	if userExists(username) {
		return errors.New("user exists")
	}else{
		client.username = username
		clientsList = append(clientsList, client)
	}
	return nil
}

/**
starts a go-routine that keeps listening to the channel 'clientChannel'.
Whenever there is a new message by a client, that client is put into this channel by  'clientHandler' function.
 */
func messageListener(clientChannel chan Client) {
	for {
		go requestHandler(<-clientChannel)
	}
}

/**
This is run for each client.
When there is a new message by a client, that client is put into the 'clientChannel' along with a message attached to her.
 */
func clientHandler(client Client, clientChannel chan Client) {

	go fmt.Fprintf(client.conn, "Connection successful"+"\n")

	//TODO put the client message into a channel, follow the same mechanism
	//keep listening to this client
	for {
		buf := bufio.NewReader(client.conn)
		request, err := buf.ReadString('\n')

		if err != nil {
			fmt.Printf("Client disconnected.\n")
			break
		}

		client.message = request
		clientChannel <- client
	}
}

func main() {

	//--------------- log setup ------------------
	f, err := os.OpenFile("server_logs", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    fmt.Printf("error opening file: %v",err)
	}
	defer func(){
		f.Close()
	}()

	log.SetOutput(f)
	//--------------- log setup ------------------

	clientChannel := make(chan Client)
	go messageListener(clientChannel)
	count := 0

	// returns a net.Listener object
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is ready.")

	/**
	when ever a new connection comes, create a new client object and start a dedicated go-routine for that client.
	 */
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Accepted connection.")

		count++
		go clientHandler(Client{count, "", conn, "", ""}, clientChannel)
	}
}
