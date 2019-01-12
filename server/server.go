package main

import (
	"bufio"
	"fmt"
	"github.com/bhargavbhegde7/GoChat/common"
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

var clientsList []Client
//var clientsList = make(map[string]Client)

func sendResponse(conn net.Conn, tag string, message string){
	user := &common.Response{ResTag: tag, Message: message}
	response, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(conn, string(response)+"\n")
}

func requestHandler(client Client){

	request := Request{}
	json.Unmarshal([]byte(client.message), &request)

	switch request.Reqtag {
	case common.GET_CLIENTS:
			clients := ""
			for _, client := range clientsList {
			    clients = clients+" : "+client.username
			}
			go sendResponse(client.conn, common.CLIENTS_LIST, clients)

		break
	case common.LOGIN:
			go sendResponse(client.conn, common.LOGIN_SUCCESS, common.NONE)

		break
	case common.SIGNUP:
			err := signup(client, request.Username)
			if err != nil{
				go sendResponse(client.conn, common.SIGNUP_FAILURE, common.NONE)
			}else{
				go sendResponse(client.conn, common.SIGNUP_SUCCESSFUL, common.NONE)
			}

		break
	case common.SELECT_TARGET:
			err := setTarget(client, request.Username)
			if err != nil{
				go sendResponse(client.conn, common.TARGET_FAIL, common.NONE)
			}else{
				go sendResponse(client.conn, common.TARGET_SET, common.NONE)
				//plus attach the public key to the json
			}

		break
		default:
			go sendResponse(client.conn, common.NONE, common.NONE)

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
