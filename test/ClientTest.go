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
	size := 50
	duration := time.Duration(5)

	for i := 0; i < size; i++ {
		pubKeyFilePath := "D:/work/gopath/src/github.com/bhargavbhegde7/GoChat/server/pub_key"
		privKeyFilePath := "D:/work/gopath/src/github.com/bhargavbhegde7/GoChat/server/priv_key"

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

	//client_utils.StartREPL(clients[0])

	time.Sleep(duration * time.Second)

	for i := 0; i < size; i++ {
		signupRequest := common.NewRequest(common.SIGNUP, clients[i].Username, clients[i].PubKey, []byte(common.NONE))
		client_utils.SendPlainTextRequest(clients[i].Conn, signupRequest)
	}

	time.Sleep(duration * time.Second)

	for i := 0; i < size; i++ {
		//clients[((i+1)%size)]
		//clients[i]
		selectTargetRequest := common.NewRequest(common.SELECT_TARGET, clients[(i+1)%size].Username, clients[i].PubKey, []byte(common.NONE))
		client_utils.SendPlainTextRequest(clients[i].Conn, selectTargetRequest)
	}

	time.Sleep(4 * time.Second)

	for i := 0; i < size; i++ {
		client_utils.ParseInput("hello", clients[i])
	}

	in := bufio.NewReader(os.Stdin)
	in.ReadString('\n')
}
