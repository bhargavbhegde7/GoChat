package main

import (
	"fmt"
	"github.com/bhargavbhegde7/GoChat/client_utils"
	"net"
)

func main() {

	items := []net.Conn{}

	for i := 0; i < 10; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			panic(err)
		}
		go listenToServer(conn)

		//common.AsymmetricPrivateKeyDecryption([]byte("ABC€"), []byte("ABC€"))

		items = append(items, conn)
	}

	//use parse input method to user input

	fmt.Println("Hello")
}
