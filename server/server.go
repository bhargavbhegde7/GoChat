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
}

const GET_CLIENTS 	= "~&#get_clients#&~"
const LOGIN 				= "~&#login#&~"
const SIGNUP 				= "~&#login#&~"
const SELECT_TARGET = "~&#selectTarget#&~"
const TARGET_FAIL 	= "~&#targetFail#&~"
const TARGET_SET 		= "~&#targetset#&~"

var clientsList []Client

func handleRequest(client Client){

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
			err := signup(client)
			if err != nil{
				go fmt.Fprintf(client.conn, "~&#signupfailure#&~"+"\n")
			}else{
				go fmt.Fprintf(client.conn, "~&#signupsuccess#&~"+"\n")
			}

		break
	case SELECT_TARGET:
			err := setTarget(client, "goodbytes")
			if err != nil{
				go fmt.Fprintf(client.conn, TARGET_FAIL+"\n")
			}else{
				go fmt.Fprintf(client.conn, TERGET_SET+"\n")
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

func getUsername(message string) string{
	return "bhegde"
}

func userExists(username string) bool{
	for _, client := range clientsList {
	    if username == client.username{
				return true
			}
	}
	return false
}

func signup(client Client) error{
	//check for existing username and
	//send either ~&#signupsuccess#&~
	//or ~&#error#&~ + ~&#signupfailure#&~
	if userExists(getUsername(client.message)) {
		return errors.New("user exists")
	}else{
		client.username = client.message
		clientsList = append(clientsList, client)
	}
	return nil
}

func handleMessage(client chan Client) {
	for {
		client := <-client
		//message := client.message
		//fmt.Printf("\nclient %d said : "+message, client.id)
		go handleRequest(client)
	}
}

func handleClient(client Client, clientChannel chan Client) {

	go fmt.Fprintf(client.conn, "Connection successful"+"\n")

	//TODO put the client message into a channel, follow the same mechanism
	//keep listening to this client
	for {
		buf := bufio.NewReader(client.conn)
		message, err := buf.ReadString('\n')

		if err != nil {
			fmt.Printf("Client disconnected.\n")
			break
		}

		clientChannel <- Client{client.id, message, client.conn, "", ""}
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
	go handleMessage(clientChannel)
	count := 0

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is ready.")

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Accepted connection.")

		count++
		go handleClient(Client{count, "", conn, "", ""}, clientChannel)
	}
}
