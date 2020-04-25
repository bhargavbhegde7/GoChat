package main

import (
	"GoChat/client_utils"
	"GoChat/common"
	"fmt"
	"net"
	"strconv"
)

func main() {

	var clients []client_utils.Client

	for i := 0; i < 1; i++ {
		pubKeyFilePath := ""
		privKeyFilePath := ""

		client := &client_utils.Client{Conn: nil, Targetpubkey: nil, Username: nil, ServerPubKey: nil, ServerKey: nil, PubKey: nil, PrivKey: nil}

		client.PubKey, client.PrivKey = common.InitRSA(pubKeyFilePath, privKeyFilePath)

		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			panic(err)
		}

		client.Conn = conn

		go client_utils.ListenToServer(client)

		clients = append(clients, client)
	}

	//use parse input method to user input

	for i := 0; i < 1; i++ {
		signupRequest := common.NewRequest(common.SIGNUP, "user0"+strconv.Itoa(i), clients[i].PubKey, []byte(common.NONE))
		client_utils.SendPlainTextRequest(clients[i].Conn, signupRequest)

		selectTargetRequest := common.NewRequest(common.SELECT_TARGET, "user0"+strconv.Itoa(i), clients[i].PubKey, []byte(common.NONE))
		client_utils.SendPlainTextRequest(clients[i].Conn, selectTargetRequest)
	}

	//client_utils.ParseInput("hello", clients[0])

	fmt.Println("Hello")
}
