package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func changeCountry(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address Address = Address{}
	changeCountry(&address)

	fmt.Println(address)
}
