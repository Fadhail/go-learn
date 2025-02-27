package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {
		if i == 7 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}
}
