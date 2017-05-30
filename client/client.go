package main

import "net"
import "fmt"
import "bufio"
import "os"
import "log"
import "github.com/fatih/color"

func listenToUser(conn net.Conn) {
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		//fmt.Print("\n>")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		log.Println("read the message from user : "+text)

		// send to socket
		log.Println("sending : "+text)
		fmt.Fprintf(conn, text+"\n")
	}
}

func main() {

	//-------------- terminal setups ------------
	//color.Set(color.BgBlue)

	//--------------- log setup ------------------
	f, err := os.OpenFile("client_logs", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    fmt.Printf("error opening file: %v",err)
	}
	defer func(){
		//color.Set(color.FgWhite)
		f.Close()
	}()

	log.SetOutput(f)
	//--------------- log setup ------------------

	// connect to this socket
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	log.Println("connected to the server : ")

	go listenToUser(conn)

	//listening to the server forever
	for {
		// listen for reply
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		log.Println("message received : "+message)

		color.Green("\n" + message+"\n")
	}
}
