package main

import "fmt"

func endApp() {
	fmt.Println("End Application")

	message := recover()
	fmt.Println("Error Message:", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Application Error")
	}
}

func main() {
	runApp(true)

	fmt.Println("App Still Running")
}
