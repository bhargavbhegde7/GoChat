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
		data = make([]byte, 1024);
	)
	fmt.Println("Initiating server... (Ctrl-C to stop)");

	lis, error := net.Listen("tcp", remote);
	defer lis.Close();
	if error != nil {
		fmt.Printf("Error creating listener: %s\n", error );
		os.Exit(1);
	}
	for {
		var response string;
		var read = true;
		con, error := lis.Accept();
		if error != nil { fmt.Printf("Error: Accepting data: %s\n", error); os.Exit(2); }
		fmt.Printf("=== New Connection received from: %s \n", con.RemoteAddr());
		for read {
			n, error := con.Read(data);
			if error== nil{
				//fmt.Println(string(data[0:n])); // Debug
				response = response + string(data[0:n]);
			}else{
				fmt.Printf("Reading data : done\n", );
				read = false;
			}
		}
		fmt.Println("Data send by client: " + response);
		con.Close();
	}
}