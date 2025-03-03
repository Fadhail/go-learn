package main

import (
	"curency_converter/curency"
	"fmt"
)

func main() {
	negara := [3]string{
		"1. Indonesia",
		"2. Amerika",
		"3. Eropa",
	}

	fmt.Println("List Negara Tersedia:")
	fmt.Println(negara)
	curency.Converter()
}
