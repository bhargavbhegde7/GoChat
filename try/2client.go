package main

import (
	"fmt";
	"net";
	"os";
)

func main() {
	var (
		host = "127.0.0.1";
		port = "9001";
		remote = host + ":" + port;
	    msg string=""
	)
	var inp string
	con, error := net.Dial("tcp",remote);
	defer con.Close();
	if error != nil { fmt.Printf("Host not found: %s\n", error ); os.Exit(1); }

	for inp!="quit" {
	    fmt.Scanf("%s",&inp);
	    msg+=inp;
	}
	in, error := con.Write([]byte(msg));
	if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }

	fmt.Println("Connection OK");

}