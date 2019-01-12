package main

import (
  "fmt"
  "net"
  "bufio"
  "github.com/fatih/color"
  "strings"
  //"encoding/json"
)

const PREFIX 			= "~&#"
const SUFFIX			= "#&~"
const GET_CLIENTS   	= PREFIX + "get_clients" + SUFFIX
const SIGNUP 	    	= PREFIX + "signup" + SUFFIX
const SELECT_TARGET 	= PREFIX + "selectTarget" + SUFFIX
const SIGNUP_SUCCESSFUL = PREFIX + "signupsuccess" + SUFFIX
const TARGET_SET    	= PREFIX + "targetset" + SUFFIX
const CONTROL_MSG    	= PREFIX + "control" + SUFFIX
const MESSAGE    	    = PREFIX + "control" + SUFFIX

var signedIn bool
var targetuser string
var targetpubkey string
var username string
var pubkey string
var targetPublicKey string

func signup(conn net.Conn){

  fmt.Printf("\nusername >>")
  _, err := fmt.Scanf("%s\n", &username)

  if err != nil {
    fmt.Println("error getting username input")
  }else{
    if strings.TrimSpace(username) != ""{
      pubkey = "abcdef"+"-"+username
      jsonString := `{"reqtag":"`+SIGNUP+`","username":"`+username+`","pubkey":"`+pubkey+`","message":"`+CONTROL_MSG+`"}`
    	fmt.Fprintf(conn, jsonString+"\n")

      //----------------------------------------
      // listen for reply
      message, err := bufio.NewReader(conn).ReadString('\n')
      if err != nil {
        panic(err)
      }

      if strings.Contains(message, SIGNUP_SUCCESSFUL){
        color.Green("sign up successful")
        signedIn = true
      }else{
        color.Red("signup error received")
      }

      //----------------------------------------
    }else{
      fmt.Println("\nerror in the username\n\n")
    }
  }
}

func getClients(conn net.Conn){
  jsonString := `{"reqtag":"`+GET_CLIENTS+`","username":"`+username+`","pubkey":"`+pubkey+`","message":"`+CONTROL_MSG+`"}`
    fmt.Fprintf(conn, jsonString+"\n")

  message, err := bufio.NewReader(conn).ReadString('\n')
  if err != nil {
    panic(err)
  }

  color.Green(message)

}

func selectTarget(conn net.Conn){
  var username string
  fmt.Printf("\ntarget username >>")
  _, err := fmt.Scanf("%s\n", &username)

  if err != nil {
    fmt.Println("error getting target username input")
  }else{
    if strings.TrimSpace(username) != ""{
      jsonString := `{"reqtag" : "`+SELECT_TARGET+`","username":"`+username+`","pubkey":"`+pubkey+`","message":"`+CONTROL_MSG+`"}`
      fmt.Fprintf(conn, jsonString+"\n")

      //----------------------------------------
      // listen for reply
      message, err := bufio.NewReader(conn).ReadString('\n')
      if err != nil {
        panic(err)
      }

      if strings.Contains(message, TARGET_SET){
        color.Green("target selection successful")
        //plus get the pub key from the json too
      }else{
        color.Red("target error received")
      }

      //----------------------------------------

      quit := make(chan int)
      go startListeningToServer(conn, quit)
      startListeningToMessages(conn, quit)

    }else{
      fmt.Println("\nerror in the username\n\n")
    }
  }
}

func startListeningToMessages(conn net.Conn, quit chan int){
  fmt.Println("starting chat . . .")
  fmt.Println("enter ^ to go to options")

  for{
    var message string
    fmt.Printf("\n>>")
    _, err := fmt.Scanf("%s\n", &message)
    if err != nil {
      fmt.Println(err)
    }else{
      if strings.Contains(message, "~~"){
        break
      }else {
        jsonString := `{"reqtag" : "`+MESSAGE+`","username":"`+username+`","pubkey":"`+pubkey+`","message":"`+message+`"}`
          fmt.Fprintf(conn, jsonString+"\n")
      }
    }
  }//infinite for loop ends
  fmt.Println("closing chat")
  close(quit)
}

func startListeningToServer(conn net.Conn, quit chan int){
  fmt.Println("starting server listener thread")

  for {
    select {
    case <-quit:
      return
    default:
      message, err := bufio.NewReader(conn).ReadString('\n')
      if err != nil {
        panic(err)
      }else {
        color.Yellow(message)
      }
    }
  }// infinite for ends
}

func parseOption(option int, conn net.Conn){
  switch option {
    case 1:
      fmt.Println(". . . login")
      break;
    case 2:
      signup(conn)
      break
    case 3:
      getClients(conn)
      break
    case 4:
      selectTarget(conn)
      break
    default:
      fmt.Println("no match found")
      break
  }
}

func main(){

  signedIn = false

  conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to the server : ")

  //----------------------------------------
  // listen for reply
  message, err := bufio.NewReader(conn).ReadString('\n')
  if err != nil {
    panic(err)
  }

  color.Red(message)

  //----------------------------------------

for{

    fmt.Println("instructions and options")

    fmt.Println("1. login")
    fmt.Println("2. sign up")
    fmt.Println("3. get users")
    fmt.Println("4. select target")
    fmt.Println("ctrl+c to exit")

    var option int
    fmt.Printf("\n>>")
    _, err := fmt.Scanf("%d\n", &option)
    if err != nil {
      fmt.Println(err)
    }else{
      parseOption(option, conn)
    }
  }//infinite for loop ends
}
