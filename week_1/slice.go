package main

import "fmt"

func main() {
	var cars = [...]string{"Car1", "Car2", "Car3", "Car4", "Car5", "Car6", "Car7", "Car8", "Car9", "Car10", "Car11", "Car12", "Car13", "Car14", "Car15"}

	var slices = cars[5:9]
	slices1 := cars[5:]
	slices2 := cars[:7]
	slices3 := cars[:]

	fmt.Println(slices[1])
	fmt.Println(slices)
	fmt.Println(slices1)
	fmt.Println(slices2)
	fmt.Println(slices3)
}
