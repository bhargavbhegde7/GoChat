package main

import (
	"GoChat/client_utils"
	"GoChat/common"
	"net"
	"os"
)

func main() {

	argsWithoutProg := os.Args[1:]

	pubKeyFilePath := argsWithoutProg[0]
	privKeyFilePath := argsWithoutProg[1]

	client_utils.PubKey, client_utils.PrivKey = common.InitRSA(pubKeyFilePath, privKeyFilePath)

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	go client_utils.ListenToServer(conn)

	//REPL
	client_utils.StartREPL(conn)
}
