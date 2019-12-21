package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bhargavbhegde7/GoChat/common"
	"log"
	"net"
	"os"
)

type Client struct {
	id       int
	message  string
	conn     net.Conn
	username string
	target   string
	pubKey   []byte
}

var clientsList []*Client
var pubKey []byte
var privKey []byte

func sendResponse(conn net.Conn, response *common.Response) {
	responseStr, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = fmt.Fprintf(conn, string(responseStr)+"\n")
	if err != nil {
		fmt.Println(err)
	}
}

func requestHandler(client *Client) {

	request := common.Request{}

	err := json.Unmarshal([]byte(client.message), &request)
	if err != nil {
		fmt.Println(err)
	}

	switch request.ReqTag {

	case common.GET_CLIENTS:
		clients := ""
		for _, eachClient := range clientsList {
			clients = clients + " : " + eachClient.username
		}
		response := common.NewResponse(common.CLIENTS_LIST, []byte(clients), common.NONE)
		go sendResponse(client.conn, response)

		break
	case common.LOGIN:
		response := common.NewResponse(common.LOGIN_SUCCESS, []byte(common.NONE), common.NONE)
		go sendResponse(client.conn, response)

		break
	case common.SIGNUP:
		err := signup(client, request.Username)
		if err != nil {
			response := common.NewResponse(common.SIGNUP_FAILURE, []byte("User already exists"), common.NONE)
			go sendResponse(client.conn, response)
		} else {
			response := common.NewResponse(common.SIGNUP_SUCCESSFUL, []byte(common.NONE), common.NONE)
			go sendResponse(client.conn, response)
		}

		break
	case common.SELECT_TARGET:
		err := setTarget(client, request.Username)
		if err != nil {
			response := common.NewResponse(common.TARGET_FAIL, []byte(common.NONE), common.NONE)
			go sendResponse(client.conn, response)
		} else {

			targetClient := getClient(client.target)

			if targetClient != nil {
				pubkeyOfTargetClient := targetClient.pubKey
				response := common.NewResponse(common.TARGET_SET, pubkeyOfTargetClient, common.NONE)
				go sendResponse(client.conn, response)
				//plus attach the public key to the json
			} else {
				response := common.NewResponse(common.TARGET_FAIL, []byte(common.NONE), common.NONE)
				go sendResponse(client.conn, response)
			}
		}

		break
	case common.CLIENT_MESSAGE:
		targetClient := getClient(client.target)

		if targetClient != nil {
			response := common.NewResponse(common.CLIENT_MESSAGE, request.Message, request.Username)
			go sendResponse(targetClient.conn, response)
		} else {
			response := common.NewResponse(common.TARGET_NOT_SET, []byte("Please set a target user"), common.NONE)
			go sendResponse(client.conn, response)
		}

		break
	case common.SERVER_KEY_EXCHANGE:
		encryptedClientKey := request.Message
		clientKey := common.AsymmetricPrivateKeyDecryption(privKey, encryptedClientKey)
		encryptedACK := common.SymmetricEncryption(clientKey, common.SERVER_KEY_ACK)

		client.pubKey = request.Pubkey
		response := common.NewResponse(common.SERVER_KEY_ACK, encryptedACK, common.NONE)
		go sendResponse(client.conn, response)
		break
	default:
		response := common.NewResponse(common.NONE, []byte(common.NONE), common.NONE)
		go sendResponse(client.conn, response)

		break
	}
}

func getClient(username string) *Client {
	var retVal *Client
	for _, client := range clientsList {
		if username == client.username {
			retVal = client
		}
	}
	return retVal
}

func setTarget(client *Client, username string) error {
	//set this username as the target for this client
	if userExists(username) {
		client.target = username
		return nil
	}
	return errors.New("user does not exist")
}

func userExists(username string) bool {
	for _, client := range clientsList {
		if username == client.username {
			return true
		}
	}
	return false
}

func signup(client *Client, username string) error {
	//check for existing username and
	//send either ~&#signupsuccess#&~
	//or ~&#error#&~ + ~&#signupfailure#&~
	if userExists(username) {
		return errors.New("user exists")
	} else {
		client.username = username
		clientsList = append(clientsList, client)
	}
	return nil
}

/**
This is run for each client.
When there is a new message by a client, that client is put into the 'clientChannel' along with a message attached to her.
*/
func clientHandler(client *Client) {

	response := common.NewResponse(common.CONNECTION_SUCCESSFUL, pubKey, common.NONE)
	go sendResponse(client.conn, response)

	//TODO put the client message into a channel, follow the same mechanism
	//keep listening to this client
	for {
		buf := bufio.NewReader(client.conn)
		request, err := buf.ReadString('\n')

		if err != nil {
			fmt.Printf("Client disconnected.\n")
			break
		}

		client.message = request
		//clientChannel <- client

		go requestHandler(client)
	}
}

func main() {
	fmt.Println("Server is ready.")
	pubKey, privKey = common.InitRSA()
	//--------------- log setup ------------------
	f, err := os.OpenFile("server_logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	log.SetOutput(f)
	//--------------- log setup ------------------

	count := 0

	// returns a net.Listener object
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is ready.")

	/**
	when ever a new connection comes, create a new client object and start a dedicated go-routine for that client.
	*/
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Accepted connection.")

		count++
		go clientHandler(&Client{count, "", conn, "", "", []byte("")})
	}
}
