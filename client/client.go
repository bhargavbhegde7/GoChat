package main

import (
	"fmt"
	"net"
)

var targetpubkey string
var username string
var pubkey string
var serverPubKey string
var serverKey string

func main(){

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to the server")

	go channelSelector()
	go listenToServer(conn)

	//REPL
	startREPL(conn)
}
