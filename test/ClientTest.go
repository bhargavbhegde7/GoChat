package main

import (
	"GoChat/client_utils"
	"GoChat/common"
	"bufio"
	"net"
	"os"
	"strconv"
)

func main() {

	var clients []client_utils.Client
	size := 10

	for i := 0; i < size; i++ {
		pubKeyFilePath := "/home/bhegde/go/src/GoChat/client/pub_key"
		privKeyFilePath := "/home/bhegde/go/src/GoChat/client/priv_key"

		client := &client_utils.Client{Conn: nil, Targetpubkey: nil, Username: "", ServerPubKey: nil, ServerKey: nil, PubKey: nil, PrivKey: nil}

		client.PubKey, client.PrivKey = common.InitRSA(pubKeyFilePath, privKeyFilePath)

		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			panic(err)
		}

		client.Conn = conn

		go client_utils.ListenToServer(client)

		clients = append(clients, *client)
	}

	//use parse input method to user input

	//sign up all users
	for i := 0; i < size; i++ {
		signupRequest := common.NewRequest(common.SIGNUP, "user0"+strconv.Itoa(i), clients[i].PubKey, []byte(common.NONE))
		client_utils.SendPlainTextRequest(clients[i].Conn, signupRequest)
	}

	//select target all users
	for i := 0; i < size; i++ {
		selectTargetRequest := common.NewRequest(common.SELECT_TARGET, "user0"+strconv.Itoa((i+1)%size), clients[i].PubKey, []byte(common.NONE))
		client_utils.SendPlainTextRequest(clients[i].Conn, selectTargetRequest)
	}

	for i := 0; i < size; i++ {
		client_utils.ParseInput("hello", &clients[i])
	}

	in := bufio.NewReader(os.Stdin)
	in.ReadString('\n')
}
