package main

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	id      int
	message string
	conn    net.Conn
}

func handleMessage(client chan Client) {
	for {
		// Wait for the next client to come off the queue.
		client := <-client

		message := client.message
		fmt.Printf("\nclient %d said : "+message, client.id)

		// Send back the response.
		go client.conn.Write([]byte("received " + client.message))
	}
}

func handleNewClient(client Client, clientChannel chan Client) {

	for {
		buf := bufio.NewReader(client.conn)
		message, err := buf.ReadString('\n')

		if err != nil {
			fmt.Printf("Client disconnected.\n")
			break
		}

		clientChannel <- Client{client.id, message, client.conn}
	}
}

func main() {
	clientChannel := make(chan Client)
	go handleMessage(clientChannel)
	count := 0

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is ready.")

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Accepted connection.")

		count++
		go handleNewClient(Client{count, "", conn}, clientChannel)
	}
}
