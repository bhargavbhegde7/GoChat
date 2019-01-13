package main

import (
	"encoding/json"
	"fmt"
	"github.com/bhargavbhegde7/GoChat/common"
	"net"
	"strings"
)

func sendRequest(conn net.Conn, request* common.Request){
	reqStr, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(conn, string(reqStr)+"\n")
}

func getClients(conn net.Conn){
	request := common.NewRequest(common.GET_CLIENTS, username, pubkey, common.NONE)
	go sendRequest(conn, request)
}

func selectTarget(conn net.Conn){
	var username string
	fmt.Printf("\ntarget username >>")
	_, err := fmt.Scanf("%s\n", &username)

	if err != nil {
		fmt.Println("error getting target username input")
	}else{
		if strings.TrimSpace(username) != ""{
			request := common.NewRequest(common.SELECT_TARGET, username, pubkey, common.NONE)
			go sendRequest(conn, request)
		}else{
			fmt.Println("\nerror in the username\n\n")
		}
	}
}

func login(conn net.Conn){

}

func sendMessage(conn net.Conn, request *common.Request){
	go sendRequest(conn, request)
}

func signup(conn net.Conn){

	fmt.Printf("\nusername >>")
	_, err := fmt.Scanf("%s\n", &username)

	if err != nil {
		fmt.Println("error getting username input")
	}else{
		if strings.TrimSpace(username) != ""{
			pubkey = "abcdef"+"-"+username

			request := common.NewRequest(common.SIGNUP, username, pubkey, common.NONE)
			go sendRequest(conn, request)
		}else{
			fmt.Println("\nerror in the username\n\n")
		}
	}
}

func generateRandomKey() string{
	return "randomKey"
}

func initServerKeyExchange(conn net.Conn){
	key := generateRandomKey()
	encryptedKey := common.AsymmetricEncryption(serverPubKey, key)

	request := common.NewRequest(common.SERVER_KEY_EXCHANGE, username, pubkey, encryptedKey)
	go sendRequest(conn, request)
}