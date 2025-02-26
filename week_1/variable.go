package main

import "fmt"

func main() {
	// Can change the value of variable
	var nama = "John"
	nama = "Doe"
	fmt.Println(nama)

	// Can't change the value of constant
	const umur int = 20
	fmt.Println(umur)
}
