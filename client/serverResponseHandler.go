package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/bhargavbhegde7/GoChat/common"
	"github.com/fatih/color"
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
					color.Red("Couldn't sign up.")
					break

				case common.CLIENT_MESSAGE:
					messageChannel<-response
					break

				case common.CONNECTION_SUCCESSFUL:
					color.Green("Connected to server")
					break

				case common.SIGNUP_SUCCESSFUL:
					color.Green("Signup was successful. You can now select a target and send messages")
					break

				case common.CLIENTS_LIST:
					color.Green(response.Message)
					break

				case common.TARGET_SET:
					color.Green("Target user is set. Target public key saved.")
					targetpubkey = response.Message
					break

				case common.TARGET_FAIL:
					color.Red("Couldn't set the target.")
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
