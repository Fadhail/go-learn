package main

import "fmt"

func getFullName() (string, string) {
	return "John", "Doe"
}

func main() {
	firstname, lastname := getFullName()
	fmt.Println("Hello", firstname, lastname)

	// Ignore the lastname
	firstname, _ = getFullName()
	fmt.Println("Hello", firstname)
}
