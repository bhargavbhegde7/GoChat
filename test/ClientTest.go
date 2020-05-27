package main

import (
	"bufio"
	"github.com/bhargavbhegde7/GoChat/client_utils"
	"github.com/bhargavbhegde7/GoChat/common"
	"net"
	"os"
	"strconv"
	"sync"
)

func main() {

	var clients []*client_utils.Client
	size := 50

	wg1 := &sync.WaitGroup{}

	//create a boatload of clients
	for i := 0; i < size; i++ {
		pubKeyFilePath := "/home/bhegde/go/src/github.com/bhargavbhegde7/GoChat/client/pub_key"
		privKeyFilePath := "/home/bhegde/go/src/github.com/bhargavbhegde7/GoChat/client/priv_key"
		wg1.Add(1)

		client := client_utils.Client{Conn: nil, Targetpubkey: nil, Username: "", ServerPubKey: nil, ServerKey: nil, PubKey: nil, PrivKey: nil}

		client.PubKey, client.PrivKey = common.InitRSA(pubKeyFilePath, privKeyFilePath)

		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			panic(err)
		}

		client.Conn = conn
		client.Username = "user0" + strconv.Itoa(i)

		go func(client *client_utils.Client) {
			client_utils.ListenToServer(client)
			wg1.Done()
		}(&client)

		clients = append(clients, &client)
	}

	wg1.Wait()

	//TODO use wait groups here instead of sleeping till all clients have started to listen to the server
	//time.Sleep(duration * time.Second)

	wg2 := &sync.WaitGroup{}
	//sign up all these clients
	for i := 0; i < size; i++ {
		wg2.Add(1)
		signupRequest := common.NewRequest(common.SIGNUP, clients[i].Username, clients[i].PubKey, []byte(common.NONE))

		go func() {
			client_utils.SendPlainTextRequest(clients[i].Conn, signupRequest)
			wg2.Done()
		}()

	}
	wg2.Wait()

	//TODO use wait groups here instead of sleeping till all clients have made a sign up request
	//time.Sleep(duration * time.Second)

	wg3 := &sync.WaitGroup{}
	//select one next username for each user
	//For example, select user01 as target for user00 and so on
	for i := 0; i < size; i++ {
		wg3.Add(1)
		selectTargetRequest := common.NewRequest(common.SELECT_TARGET, clients[(i+1)%size].Username, clients[i].PubKey, []byte(common.NONE))

		go func() {
			client_utils.SendPlainTextRequest(clients[i].Conn, selectTargetRequest)
			wg3.Done()
		}()

	}
	wg3.Wait()

	//TODO use wait groups instead of waiting with sleep till target selection is over.
	//time.Sleep(duration * time.Second)

	wg4 := &sync.WaitGroup{}
	//make every client send out a message to its target
	for i := 0; i < size; i++ {
		wg4.Add(1)

		go func() {
			client_utils.ParseInput("hello", clients[i])
			wg4.Done()
		}()
	}
	wg4.Wait()

	//wait for enter key
	in := bufio.NewReader(os.Stdin)
	in.ReadString('\n')
}
