package main

import "fmt"

type Students struct {
	FirstName string
	LastName  string
	Npm       int
	Age       int
}

func (students Students) showBio() {
	fmt.Println("Here we go this is your data!")
	fmt.Println("Your name :", students.FirstName, students.LastName)
	fmt.Println("Yor NPM :", students.Npm)
	fmt.Println("Your Age :", students.Age)
}

func getStudent() (string, string, int, int) {
	fmt.Print("First Name : ")
	first := ""
	fmt.Scanln(&first)

	fmt.Print("Last Name : ")
	last := ""
	fmt.Scanln(&last)

	fmt.Print("NPM :")
	npm := 0
	fmt.Scanln(&npm)

	fmt.Print("Age :")
	age := 0
	fmt.Scanln(&age)

	return first, last, npm, age
}

func main() {
	fName, lName, npm, age := getStudent()
	name := Students{fName, lName, npm, age}
	name.showBio()
}
