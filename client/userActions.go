package main

import (
	"encoding/json"
	"fmt"
	"github.com/bhargavbhegde7/GoChat/common"
	"net"
	"strings"
)

func sendPlainTextRequest(conn net.Conn, request *common.Request) {
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
	request := common.NewRequest(common.GET_CLIENTS, username, pubKey, []byte(common.NONE))
	go sendPlainTextRequest(conn, request)
}

func selectTarget(conn net.Conn) {
	var username string
	fmt.Printf("\ntarget username >>")
	_, err := fmt.Scanf("%s\n", &username)

	if err != nil {
		fmt.Println("error getting target username input")
	} else {
		if strings.TrimSpace(username) != "" {
			request := common.NewRequest(common.SELECT_TARGET, username, pubKey, []byte(common.NONE))
			go sendPlainTextRequest(conn, request)
		} else {
			fmt.Println("\nerror in the username\n\n")
		}
	}
}

func login(conn net.Conn) {

}

func sendMessage(conn net.Conn, request *common.Request) {
	request.Message = common.AsymmetricPublicKeyEncryption(targetpubkey, request.Message)
	go sendPlainTextRequest(conn, request)
}

func signup(conn net.Conn) {

	fmt.Printf("\nusername >>")
	_, err := fmt.Scanf("%s\n", &username)

	if err != nil {
		fmt.Println("error getting username input")
	} else {
		if strings.TrimSpace(username) != "" {
			//pubkey = "abcdef" + "-" + username

			request := common.NewRequest(common.SIGNUP, username, pubKey, []byte(common.NONE))
			go sendPlainTextRequest(conn, request)
		} else {
			fmt.Println("\nerror in the username\n\n")
		}
	}
}

func initServerKeyExchange(conn net.Conn) {
	serverKey = common.GenerateRandomKey()
	encryptedKey := common.AsymmetricPublicKeyEncryption(serverPubKey, serverKey)

	request := common.NewRequest(common.SERVER_KEY_EXCHANGE, username, pubKey, encryptedKey)
	go sendPlainTextRequest(conn, request)
}
