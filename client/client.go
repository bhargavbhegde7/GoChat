package main

import "net"
import "fmt"
import "bufio"
import "os"
import "log"
import "github.com/fatih/color"
import "strings"
//import "errors"

func isControlMsg(msg string) bool{
	if string(msg[0]) == "^"{
		return true
	}else{
		return false
	}
}

func handleInput(text string, conn net.Conn){

	if isControlMsg(text){
		fmt.Println("\n control message found\n")

		//take some more data and stuff.....

		//reconfig
	}else{
		fmt.Println("\ndata message found\n")

		//send to the target client

		log.Println("sending : "+text)
		fmt.Fprintf(conn, text+"\n")
	}
}

func listenToUser(conn net.Conn) {
	for {

		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		log.Println("read the message from user : "+text)

		if strings.TrimSpace(text) != ""{
			handleInput(text, conn)
		}else{
			color.Red("\nerror in the text\n\n")
		}
	}//infinite for loop ends
}

func login(){

}

func signup(conn net.Conn){


	//jsonString := "{'type' : '~&#signup#&~','username' : '"+username+"','pubkey' : 'abcdef'}"
	//fmt.Fprintf(conn, jsonString+"\n")
}

func config(conn net.Conn){
	for {
		color.Yellow("\ninstructions and options")
		fmt.Println("1. login")
		fmt.Println("2. sign up")
		fmt.Println("ctrl+c to exit")

		var option int
    _, err := fmt.Scanf("%d", &option)
		if err != nil {
			panic(err)
			color.Red("select an option")
		}else{
			if option == 1{
				color.Yellow("\nlogging you in\n")
				login()
				break
			}else if option == 2{
				color.Yellow("\ncalling signup\n")
				//signup(conn)
				fmt.Print("Enter text: ")
			  var input string
			  fmt.Scanln(&input)
			  fmt.Print(input)
				//break
			}else{
				color.Red("\nselect an option\n")
			}
		}
	}//infinite for loop ends

	//return errors.New("error in config")
	//return nil
}

func isError(msg string) bool{
	if strings.Contains(msg, "~&#error#&~"){
		return true
	}else{
		return false
	}
}

func handleServerMessage(chanMsg chan string){
	for {
		// Wait for the next server message to come off the queue.
		message := <-chanMsg

		if isError(message) {
			color.Green("\n" + message+"\n")
		}else if strings.Contains(message, "~&#clients#&~") {
			color.Green("\n" + "got all the cliends details"+"\n")
		}else if strings.Contains(message, "~&#target_disconnected#&~") {
			color.Green("\n" + "remote client disconnected"+"\n")
		}else if strings.Contains(message, "~&#signupsuccess#&~") {
			color.Green("\n" + "sign up successful"+"\n")
		}else if strings.Contains(message, "~&#signupfailure#&~") {
			color.Green("\n" + "sign up failed"+"\n")
		}else{
			color.Green("\n" + message+"\n")
		}
	}
}

func main() {

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

	//initial hand shakes with the server
	config(conn)

	//go listenToUser(conn)

	/*
	serverMessageChannel := make(chan string)
	go handleServerMessage(serverMessageChannel)
	//listening to the server forever
	for {
		// listen for reply
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		log.Println("message received : "+message)

		//color.Green("\n" + message+"\n")

		serverMessageChannel <- message
	}
	*/

}
