package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"log"
)

type Client struct {
	id      int
	message string
	conn    net.Conn
}

func sendMessage(conn net.Conn, client Client){
	_, _ = conn.Write([]byte(client.message))
	// if sdf != nil {
	// 	fmt.Println("hello")
	// }

}

func handleMessage(client chan Client) {
	for {
		// Wait for the next client to come off the queue.
		client := <-client

		message := client.message
		fmt.Printf("\nclient %d said : "+message, client.id)

		// Send back the response.
		//a := "hello"
		go client.conn.Write([]byte(client.message))
		//go sendMessage(client.conn, client)
		//go client.conn.Write([]byte(a))
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

	//--------------- log setup ------------------
	f, err := os.OpenFile("server_logs", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    fmt.Printf("error opening file: %v",err)
	}
	defer func(){
		//color.Set(color.FgWhite)
		f.Close()
	}()

	log.SetOutput(f)
	//--------------- log setup ------------------

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
