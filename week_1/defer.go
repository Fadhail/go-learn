package main

import "fmt"

func logging() {
	fmt.Println("End Application")
}

func runApplication() {
	defer logging()

	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
