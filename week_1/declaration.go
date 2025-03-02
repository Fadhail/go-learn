package main

import "fmt"

func main() {
	type npm string

	var npm1 npm = "111111"

	var npm2 string = "222222"
	var contohNpm npm = npm(npm2)

	fmt.Println(npm1)
	fmt.Println(contohNpm)
}
