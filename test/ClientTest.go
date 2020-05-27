package main

import (
	"bufio"
	"github.com/bhargavbhegde7/GoChat/client_utils"
	"github.com/bhargavbhegde7/GoChat/common"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {

	var clients []*client_utils.Client
	size := 500
	duration := time.Duration(10)

	//create a boatload of clients
	for i := 0; i < size; i++ {
		pubKeyFilePath := "/home/bhegde/go/src/github.com/bhargavbhegde7/GoChat/client/pub_key"
		privKeyFilePath := "/home/bhegde/go/src/github.com/bhargavbhegde7/GoChat/client/priv_key"

		client := client_utils.Client{Conn: nil, Targetpubkey: nil, Username: "", ServerPubKey: nil, ServerKey: nil, PubKey: nil, PrivKey: nil}

		client.PubKey, client.PrivKey = common.InitRSA(pubKeyFilePath, privKeyFilePath)

		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			panic(err)
		}

		client.Conn = conn
		client.Username = "user0" + strconv.Itoa(i)

		go client_utils.ListenToServer(&client)

		clients = append(clients, &client)
	}

	time.Sleep(duration * time.Second)

	//sign up all these clients
	for i := 0; i < size; i++ {
		signupRequest := common.NewRequest(common.SIGNUP, clients[i].Username, clients[i].PubKey, []byte(common.NONE))

		go client_utils.SendPlainTextRequest(clients[i].Conn, signupRequest)

	}

	time.Sleep(duration * time.Second)

	//select one next username for each user
	//For example, select user01 as target for user00 and so on
	for i := 0; i < size; i++ {
		selectTargetRequest := common.NewRequest(common.SELECT_TARGET, clients[(i+1)%size].Username, clients[i].PubKey, []byte(common.NONE))

		go client_utils.SendPlainTextRequest(clients[i].Conn, selectTargetRequest)

	}

	time.Sleep(duration * time.Second)

	//make every client send out a message to its target
	for i := 0; i < size; i++ {

		go client_utils.ParseInput("hello", clients[i])

	}

	//wait for enter key
	in := bufio.NewReader(os.Stdin)
	in.ReadString('\n')
}
