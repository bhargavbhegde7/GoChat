package main

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	count int
	conn  net.Conn
}

func handleNewConnection(c Client) {

	for {
		buf := bufio.NewReader(c.conn)
		message, err := buf.ReadString('\n')

		if err != nil {
			fmt.Printf("Client disconnected.\n")
			break
		}

		fmt.Printf("\nclient %d said : "+message, c.count)

		go c.conn.Write([]byte("OKAY"))
		// Send back the response.
		// _, err = c.conn.Write([]byte("OKAY"))
		// if err != nil {
		// 	panic(err)
		// }

		//fmt.Fprintf(c.conn, "OKAY")

	}
}

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is ready.")

	count := 0

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Accepted connection.")

		count++
		newClient := Client{count: count, conn: conn}
		go handleNewConnection(newClient)
	}
}
