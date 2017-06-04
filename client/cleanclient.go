package main

import (
  "fmt"
  "net"
  "bufio"
  "github.com/fatih/color"
  "strings"
  //"encoding/json"
)

var signedIn bool
var targetuser string
var targetpubkey string

func signup(conn net.Conn){

  var username string
  fmt.Printf("\nusername >>")
  _, err := fmt.Scanf("%s\n", &username)

  if err != nil {
    fmt.Println("error getting username input")
  }else{
    if strings.TrimSpace(username) != ""{
      jsonString := `{"reqtag":"~&#signup#&~","username":"`+username+`","pubkey":"abcdef"}`
    	fmt.Fprintf(conn, jsonString+"\n")

      //----------------------------------------
      // listen for reply
      message, err := bufio.NewReader(conn).ReadString('\n')
      if err != nil {
        panic(err)
      }

      color.Red(message)

      if strings.Contains(message, "~&#signupsuccess#&~"){
        fmt.Println("sign up successful")
        signedIn = true
      }else{
        fmt.Println("signup error received")
      }

      //----------------------------------------
    }else{
      fmt.Println("\nerror in the username\n\n")
    }
  }
}

func getClients(conn net.Conn){
  fmt.Fprintf(conn, "~&#get_clients#&~"+"\n")

  message, err := bufio.NewReader(conn).ReadString('\n')
  if err != nil {
    panic(err)
  }

  color.Red(message)

}

func selectTarget(conn net.Conn){
  var username string
  fmt.Printf("\ntarget username >>")
  _, err := fmt.Scanf("%s\n", &username)

  if err != nil {
    fmt.Println("error getting target username input")
  }else{
    if strings.TrimSpace(username) != ""{
      jsonString := `{"reqtag" : "~&#selectTarget#&~", "username" : "`+username+`"}`
      fmt.Fprintf(conn, jsonString+"\n")

      //----------------------------------------
      // listen for reply
      message, err := bufio.NewReader(conn).ReadString('\n')
      if err != nil {
        panic(err)
      }

      color.Red(message)

      if strings.Contains(message, "~&#targetset#&~"){
        fmt.Println("target selection successful")
        //plus get the pub key from the json too
      }else{
        fmt.Println("target error received")
      }

      //----------------------------------------

    }else{
      fmt.Println("\nerror in the username\n\n")
    }
  }
}

func startListeningToUser(conn net.Conn){
  fmt.Println("starting chat . . .")
  fmt.Println("enter ^ to go to options")
  //stopSignal := make(chan bool)
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
      startListeningToUser(conn)
      break
    default:
      fmt.Println(". . .")
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
