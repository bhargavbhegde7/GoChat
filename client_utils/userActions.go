package client_utils

import (
	"GoChat/common"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

func SendPlainTextRequest(conn net.Conn, request *common.Request) {
	reqStr, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = fmt.Fprintf(conn, string(reqStr)+"\n")
	if err != nil {
		fmt.Println(err)
	}
}

func sendSymmetricEncryptedRequest(conn net.Conn, request *common.Request) {
	reqStr, err := json.Marshal(request)

	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = fmt.Fprintf(conn, string(reqStr)+"\n")
	if err != nil {
		fmt.Println(err)
	}
}

func getClients(client *Client) {
	request := common.NewRequest(common.GET_CLIENTS, client.Username, client.PubKey, []byte(common.NONE))
	go SendPlainTextRequest(client.Conn, request)
}

func selectTarget(client *Client) {
	var username string
	fmt.Printf("target username >>\n")
	_, err := fmt.Scanf("%s\n", &username)

	if err != nil {
		fmt.Println("error getting target username input")
	} else {
		if strings.TrimSpace(username) != "" {
			request := common.NewRequest(common.SELECT_TARGET, username, client.PubKey, []byte(common.NONE))
			go SendPlainTextRequest(client.Conn, request)
		} else {
			fmt.Println("\nerror in the username\n\n")
		}
	}
}

func login(conn net.Conn) {

}

func sendMessage(client *Client, request *common.Request) {
	request.Message = common.AsymmetricPublicKeyEncryption(client.Targetpubkey, request.Message)
	go SendPlainTextRequest(client.Conn, request)
}

func signup(client *Client) {

	fmt.Printf("username >>\n")
	_, err := fmt.Scanf("%s\n", &client.Username)

	if err != nil {
		fmt.Println("error getting username input")
	} else {
		if strings.TrimSpace(client.Username) != "" {
			//pubkey = "abcdef" + "-" + username

			request := common.NewRequest(common.SIGNUP, client.Username, client.PubKey, []byte(common.NONE))
			go SendPlainTextRequest(client.Conn, request)
		} else {
			fmt.Println("\nerror in the username\n\n")
		}
	}
}

func initServerKeyExchange(client *Client) {
	client.ServerKey = common.GenerateRandomKey()
	encryptedKey := common.AsymmetricPublicKeyEncryption(client.ServerPubKey, client.ServerKey)

	request := common.NewRequest(common.SERVER_KEY_EXCHANGE, client.Username, client.PubKey, encryptedKey)
	go SendPlainTextRequest(client.Conn, request)
}
