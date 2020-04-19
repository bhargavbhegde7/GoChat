package main

import (
	"os/exec"
)

func main() {
	//exec.Command("go", "build", "-o", "/home/bhegde/go/src/GoChat/client/abcd", "/home/bhegde/go/src/GoChat/client/")
	exec.Command("echo", "'build'", ">", "/home/bhegde/go/src/GoChat/client/abcd")

	/*serverProcess := exec.Command("./client/abcd", "client/pub_key", "client/priv_key")

	serverProcess.Stdout = os.Stdout
	serverProcess.Stderr = os.Stderr

	fmt.Println("SERVER TEST START") //for debug

	err := serverProcess.Start()
	if err != nil { //Use start, not run
		fmt.Println("An error occured in server: ", err) //replace with logger, or anything you want
	}

	serverProcess.Wait()
	fmt.Println("END") //for debug*/
}