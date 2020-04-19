package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	/*exec.Command("go", "run", "server/server.go")
	exec.Command("go", "build", "./client")

	clientProcess := exec.Command("client.exe")

	stdin, err := clientProcess.StdinPipe()
	if err != nil {
		fmt.Println(err) //replace with logger, or anything you want
	}
	defer stdin.Close() // the doc says clientProcess.Wait will close it, but I'm not sure, so I kept this line

	clientProcess.Stdout = os.Stdout
	clientProcess.Stderr = os.Stderr

	fmt.Println("START") //for debug
	if err = clientProcess.Start(); err != nil { //Use start, not run
		fmt.Println("An error occured: ", err) //replace with logger, or anything you want
	}

	io.WriteString(stdin, "4\n")
	clientProcess.Wait()
	fmt.Println("END") //for debug*/

	serverProcess := exec.Command("go", "run", "server/server.go", "server/pub_key", "server/priv_key")

	serverProcess.Stdout = os.Stdout
	serverProcess.Stderr = os.Stderr

	fmt.Println("SERVER TEST START") //for debug

	err := serverProcess.Start()
	if err != nil { //Use start, not run
		fmt.Println("An error occured in server: ", err) //replace with logger, or anything you want
	}

	serverProcess.Wait()
	fmt.Println("END") //for debug
}