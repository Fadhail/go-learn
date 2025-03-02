package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	// Whitout Asterik
	// var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	// var address2 *Address = &address1 // Pointer

	// address2.City = "Cimahi"
	// fmt.Println(address1)
	// fmt.Println(address2)

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// fmt.Println(address1)
	// fmt.Println(address2)

	// Asterik
	var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Cimahi"
	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
