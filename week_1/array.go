package main

import "fmt"

func main() {
	var cars [5]string

	cars[0] = "Toyota"
	cars[1] = "Honda"
	cars[2] = "Suzuki"
	cars[3] = "Mitsubishi"
	cars[4] = "Daihatsu"

	fmt.Println(cars)

	// Or
	var cars1 = [5]string{"Toyota", "Honda", "Suzuki", "Mitsubishi", "Daihatsu"}

	fmt.Println(cars1)

	// Access by index
	fmt.Println(cars1[1])

	// Change value by index
	cars1[1] = "Yamaha"

	fmt.Println(cars1)
}
