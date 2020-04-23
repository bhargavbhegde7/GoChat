package client_utils

import (
	"GoChat/client"
	"bufio"
	"fmt"
	"github.com/bhargavbhegde7/GoChat/common"
	"github.com/fatih/color"
	"net"
	"os"
	"strings"
)

func startREPL(conn net.Conn) {
	fmt.Println("enter '" + HELP + "' for instructions")
	in := bufio.NewReader(os.Stdin)

	for {
		line, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		} else {
			line = strings.TrimRight(line, "\r\n")
			parseInput(line, conn)
		}
	} //infinite for loop ends
}

func printInstructions() {
	fmt.Println("enter '" + LOGIN + "' to login")
	fmt.Println("enter '" + SIGNUP + "' to signing up with a new username")
	fmt.Println("enter '" + USERS + "' to get a list of all existing user names")
	fmt.Println("enter '" + SELECT + "' to select a user by username")
}

func parseInput(input string, conn net.Conn) {
	//fmt.Println("input : "+">>>"+input+"<<<")
	switch input {
	case HELP:
		printInstructions()
		break
	case LOGIN:
		login(conn)
		break
	case SIGNUP:
		signup(conn)
		break
	case USERS:
		getClients(conn)
		break
	case SELECT:
		selectTarget(conn)
		break
	default:
		// consider this as a message payload
		if main.targetpubkey == nil {
			color.Red("Target client is not set")
			break
		}
		request := common.NewRequest(common.CLIENT_MESSAGE, main.username, main.pubKey, []byte(input))
		sendMessage(conn, request)
		break
	}
}
