package main

import "fmt"

func getHello(name string) string {
	return "Hello, " + name
}

func main() {
	contoh1 := getHello("John")
	contoh2 := getHello("Doe")

	fmt.Println(contoh1)
	fmt.Println(contoh2)

	// Or
	contoh3 := getHello
	contoh4 := getHello

	fmt.Println(contoh3("John"))
	fmt.Println(contoh4("Doe"))
}
