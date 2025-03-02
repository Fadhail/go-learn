package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Maried() {
	man.Name = "Mr. " + man.Name
}

func main() {
	Fadel := Man{"Fadel"}
	Fadel.Maried()

	fmt.Println(Fadel)
}
