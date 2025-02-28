package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func main() {
	var john Customer
	john.Name = "John Doe"
	john.Address = "Jakarta"
	john.Age = 35

	fmt.Println(john)

	// Or
	jane := Customer{
		Name:    "Jane Doe",
		Address: "Jakarta",
		Age:     25,
	}

	fmt.Println(jane)

	// Or
	jack := Customer{"Jack Doe", "Jakarta", 30}
	fmt.Println(jack)
}
