package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/bhargavbhegde7/GoChat/common"
	"net"
)

func listenToServer(conn net.Conn){

	for {
		serverMessage, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}else {
			response := common.Response{}
			json.Unmarshal([]byte(serverMessage), &response)
			// TODO use a hashmap to save all the channels. select them based on tag and send the response to that channel
			switch response.ResTag {

			case common.SIGNUP_FAILURE:
				errorChannel<-response
				break

			case common.CLIENT_MESSAGE:
				messageChannel<-response
				break

			case common.CONNECTION_SUCCESSFUL:
				fmt.Println(string(serverMessage))
				break

			case common.SIGNUP_SUCCESSFUL:
				fmt.Println(string(serverMessage))
				break

			case common.CLIENTS_LIST:
				fmt.Println(string(serverMessage))
				break

			case common.TARGET_SET:
				// payload message will be the public key of the selected target client
				fmt.Println(string(serverMessage))
				break

			case common.TARGET_FAIL:
				fmt.Println(string(serverMessage))
				break

			case common.NONE:
				fmt.Println("Request tag did not match to any in server")
				break

			default:
				fmt.Println("unrecognised tag : "+response.ResTag+", message : "+response.Message)
			}// response tag switch ends
		}
	}// infinite for ends
}
