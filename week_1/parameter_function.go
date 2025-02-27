package main

import "fmt"

func getNameFilter(name string, filter func(string) string) {
	filterName := filter(name)
	fmt.Println("Hello, " + filterName)
}

func spamFilter(name string) string {
	if name == "spam" {
		return "..."
	} else {
		return name
	}
}

func main() {
	getNameFilter("John", spamFilter)
	getNameFilter("Doe", spamFilter)
	getNameFilter("spam", spamFilter)
}
