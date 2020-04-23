package main

import (
	"github.com/bhargavbhegde7/GoChat/client_utils"
	"github.com/bhargavbhegde7/GoChat/common"
	"net"
	"os"
)

func main() {

	argsWithoutProg := os.Args[1:]

	pubKeyFilePath := argsWithoutProg[0]
	privKeyFilePath := argsWithoutProg[1]

	pubKey, privKey = common.InitRSA(pubKeyFilePath, privKeyFilePath)

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	go client_utils.ListenToServer(conn)

	//REPL
	client_utils.StartREPL(conn)
}
