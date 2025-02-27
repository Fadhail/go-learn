package main

import "fmt"

func getName() (firstname, middlename, lastname string) {
	firstname = "John"
	middlename = "Doe"
	lastname = "Smith"

	return firstname, middlename, lastname
}

func main() {
	firstname, middlename, lastname := getName()
	fmt.Println("Hello", firstname, middlename, lastname)

	// Ignore the middlename
	firstname, _, lastname = getName()
	fmt.Println("Hello", firstname, lastname)
}
