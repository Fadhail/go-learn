package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	// Value by Copy
	// address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	// address2 := address1

	// address2.City = "Cimahi"

	// fmt.Println(address1)
	// fmt.Println(address2)

	// Pointer (Value By Reference)
	var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // Pointer

	address2.City = "Cimahi"

	fmt.Println(address1)
	fmt.Println(address2)
}
