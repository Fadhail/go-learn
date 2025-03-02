package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat1.City = "Selangor"

	alamat2.Country = "Malaysia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
