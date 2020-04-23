package main

import "fmt"

func main() {

	var first string
	var second string

	fmt.Println("Enter First String:")
	fmt.Scanf("%s\n", &first)
	fmt.Println("Entered String: \n", first)

	fmt.Println("Enter Second String:")
	fmt.Scanf("%s\n", &second)
	fmt.Println("Entered String: \n", second)
}
