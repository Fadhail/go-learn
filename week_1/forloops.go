package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("Pembatas")

	// Or
	for count := 1; count <= 10; count++ {
		fmt.Println(count)
	}

	// With range
	cars := []string{"Toyota", "Honda", "Suzuki", "Mitsubishi"}
	for index, car := range cars {
		fmt.Println(index, car)
	}
}
