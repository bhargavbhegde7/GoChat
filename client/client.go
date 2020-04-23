package main

import (
	"GoChat/client_utils"
	"github.com/bhargavbhegde7/GoChat/client_utils"
	"github.com/bhargavbhegde7/GoChat/common"
	"net"
	"os"
)

var targetpubkey []byte
var username string
var serverPubKey []byte
var serverKey []byte
var pubKey []byte
var privKey []byte

func main() {

	argsWithoutProg := os.Args[1:]

	pubKeyFilePath := argsWithoutProg[0]
	privKeyFilePath := argsWithoutProg[1]

	pubKey, privKey = common.InitRSA(pubKeyFilePath, privKeyFilePath)

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	go client_utils.listenToServer(conn)

	//REPL
	client_utils.startREPL(conn)
}
