package client_utils

import (
	"bufio"
	"fmt"
	"github.com/bhargavbhegde7/GoChat/common"
	"github.com/fatih/color"
	"os"
	"strings"
)

func StartREPL(client *Client) {
	fmt.Println("enter '" + HELP + "' for instructions")
	in := bufio.NewReader(os.Stdin)

	for {
		line, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		} else {
			line = strings.TrimRight(line, "\r\n")
			ParseInput(line, client)
		}
	} //infinite for loop ends
}

func printInstructions() {
	fmt.Println("enter '" + LOGIN + "' to login")
	fmt.Println("enter '" + SIGNUP + "' to signing up with a new username")
	fmt.Println("enter '" + USERS + "' to get a list of all existing user names")
	fmt.Println("enter '" + SELECT + "' to select a user by username")
}

func ParseInput(input string, client *Client) {
	//fmt.Println("input : "+">>>"+input+"<<<")
	switch input {
	case HELP:
		printInstructions()
		break
	case LOGIN:
		login(client.Conn)
		break
	case SIGNUP:
		signup(client)
		break
	case USERS:
		getClients(client)
		break
	case SELECT:
		selectTarget(client)
		break
	default:
		// consider this as a message payload
		if client.Targetpubkey == nil {
			color.Red("Target client is not set for client with user name " + client.Username)
			break
		}
		request := common.NewRequest(common.CLIENT_MESSAGE, client.Username, client.PubKey, []byte(input))
		sendMessage(client, request)
		break
	}
}
