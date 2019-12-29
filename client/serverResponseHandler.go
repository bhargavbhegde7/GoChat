package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/bhargavbhegde7/GoChat/common"
	"github.com/fatih/color"
	"net"
)

func listenToServer(conn net.Conn) {

	for {
		serverMessage, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		} else {
			response := common.Response{}
			json.Unmarshal([]byte(serverMessage), &response)
			// TODO use a hashmap to save all the channels. select them based on tag and send the response to that channel
			switch response.ResTag {

			case common.SIGNUP_FAILURE:
				color.Red("Couldn't sign up. reason : " + string(response.Message))
				break

			case common.CLIENT_MESSAGE:
				go messageHandler(response)
				break

			case common.CONNECTION_SUCCESSFUL:
				color.Green("Connected to server")
				serverPubKey = response.Message
				initServerKeyExchange(conn)
				break

			case common.SIGNUP_SUCCESSFUL:
				color.Green("Signup was successful. You can now select a target and send messages")
				break

			case common.CLIENTS_LIST:
				color.Green(string(response.Message))
				break

			case common.TARGET_SET:
				color.Green("Target user is set. Target public key saved.")
				targetpubkey = response.Message
				break

			case common.TARGET_FAIL:
				color.Red("Couldn't set the target.")
				break

			case common.SERVER_KEY_ACK:
				encryptedACK := response.Message
				decryptedACK := common.SymmetricDecryption(serverKey, encryptedACK)
				if common.SERVER_KEY_ACK == decryptedACK {
					color.Green("Symmetric Key exchange successful")
					//TODO send a ready message to the server so that server can understand to look for an encrypted message from now on.
					//TODO also maintain a flag in the client side that will indicate if the secure pipeline is up or not.
					//Reject all requests if this flag is not set. This should happen both on client as well as server.
				} else {
					color.Red("Symmetric Key exchange failed")
				}
				break

			case common.TARGET_NOT_SET:
				color.Red("Target user is not set. Please see instructions by inputting '~~'")
				break

			case common.NONE:
				fmt.Println("Request tag did not match to any in server")
				break

			default:
				fmt.Println("received unrecognised tag : " + response.ResTag + ", message : " + string(response.Message))
			} // response tag switch ends
		}
	} // infinite for ends
}

func messageHandler(messageResponse common.Response) {
	message := common.AsymmetricPrivateKeyDecryption(privKey, messageResponse.Message)
	color.Yellow(messageResponse.Username + " : " + message)
}
