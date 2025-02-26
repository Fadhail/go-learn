package main

import "fmt"

func main() {
	var huruf string
	huruf = "Abc"
	fmt.Println(huruf)

	// Or
	var huruf1 = "Abc"
	fmt.Println(huruf)

	// Length
	fmt.Println(len(huruf1))

	// Access by index
	fmt.Println(huruf1[1])
}
