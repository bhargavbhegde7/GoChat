package main

import (
	"fmt"
	"github.com/bhargavbhegde7/GoChat/common"
	"net"
)

var targetpubkey []byte
var username string
var serverPubKey []byte
var serverKey []byte
var pubKey []byte
var privKey []byte

func main() {

	pubKey, privKey = common.InitRSA()

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
