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

func getClients(conn net.Conn) {
	request := common.NewRequest(common.GET_CLIENTS, username, PubKey, []byte(common.NONE))
	go SendPlainTextRequest(conn, request)
}

func selectTarget(conn net.Conn) {
	var username string
	fmt.Printf("target username >>\n")
	_, err := fmt.Scanf("%s\n", &username)

	if err != nil {
		fmt.Println("error getting target username input")
	} else {
		if strings.TrimSpace(username) != "" {
			request := common.NewRequest(common.SELECT_TARGET, username, PubKey, []byte(common.NONE))
			go SendPlainTextRequest(conn, request)
		} else {
			fmt.Println("\nerror in the username\n\n")
		}
	}
}

func login(conn net.Conn) {

}

func sendMessage(conn net.Conn, request *common.Request) {
	request.Message = common.AsymmetricPublicKeyEncryption(targetpubkey, request.Message)
	go SendPlainTextRequest(conn, request)
}

func signup(conn net.Conn) {

	fmt.Printf("username >>\n")
	_, err := fmt.Scanf("%s\n", &username)

	if err != nil {
		fmt.Println("error getting username input")
	} else {
		if strings.TrimSpace(username) != "" {
			//pubkey = "abcdef" + "-" + username

			request := common.NewRequest(common.SIGNUP, username, PubKey, []byte(common.NONE))
			go SendPlainTextRequest(conn, request)
		} else {
			fmt.Println("\nerror in the username\n\n")
		}
	}
}

func initServerKeyExchange(conn net.Conn) {
	serverKey = common.GenerateRandomKey()
	encryptedKey := common.AsymmetricPublicKeyEncryption(serverPubKey, serverKey)

	request := common.NewRequest(common.SERVER_KEY_EXCHANGE, username, PubKey, encryptedKey)
	go SendPlainTextRequest(conn, request)
}
