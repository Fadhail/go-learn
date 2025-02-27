package main

import "fmt"

func sayHelloTo(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}

func main() {
	sayHelloTo("John", "Doe")
	sayHelloTo("Jane", "Doe")
	sayHelloTo("Budi", "Setiawan")
}
