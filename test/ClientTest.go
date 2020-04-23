package main

import (
	"GoChat/client_utils"
	"GoChat/common"
	"fmt"
	"net"
	"strconv"
)

func main() {

	var connections []net.Conn

	for i := 0; i < 10; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			panic(err)
		}
		go client_utils.ListenToServer(conn)

		connections = append(connections, conn)
	}

	//use parse input method to user input

	for i := 0; i < 10; i++ {
		request := common.NewRequest(common.SIGNUP, "user"+strconv.Itoa(i), client_utils.PubKey, []byte(common.NONE))
		go client_utils.SendPlainTextRequest(connections[i], request)

	}

	client_utils.StartREPL(connections[0])

	fmt.Println("Hello")
}
