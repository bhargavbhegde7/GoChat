package main

import (
	"fmt"
	"github.com/bhargavbhegde7/GoChat/common"
	"net"
)

func startREPL(conn net.Conn){
	fmt.Println("enter '"+HELP+"' for instructions")
	for{
		var input string
		fmt.Printf("\n>>")
		_, err := fmt.Scanf("%s\n", &input)
		if err != nil {
			fmt.Println(err)
		}else{
			parseInput(input, conn)
		}
	}//infinite for loop ends
}

func printInstructions(){
	fmt.Println("enter '"+LOGIN +"' to login")
	fmt.Println("enter '"+SIGNUP+"' to signing up with a new username")
	fmt.Println("enter '"+USERS +"' to get a list of all existing user names")
	fmt.Println("enter '"+SELECT+"' to select a user by username")
}

func parseInput(input string, conn net.Conn){
	switch input {
	case HELP:
		printInstructions()
		break
	case LOGIN:
		login(conn)
		break
	case SIGNUP:
		signup(conn)
		break
	case USERS:
		getClients(conn)
		break
	case SELECT:
		selectTarget(conn)
		break
	default:
		// consider this as a message payload
		request := common.NewRequest(common.CLIENT_MESSAGE, username, pubkey, input)
		go sendRequest(conn, request)
		break
	}
}