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

	client := &client_utils.Client{Conn: nil, Targetpubkey: nil, Username: "", ServerPubKey: nil, ServerKey: nil, PubKey: nil, PrivKey: nil}

	client.PubKey, client.PrivKey = common.InitRSA(pubKeyFilePath, privKeyFilePath)

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	client.Conn = conn

	go client_utils.ListenToServer(client)

	//REPL
	client_utils.StartREPL(client)
}
