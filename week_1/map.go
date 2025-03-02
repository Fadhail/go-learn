package main

import "fmt"

func main() {
	students := map[string]string{
		"nama":    "John Doe",
		"npm":     "111111",
		"jurusan": "Teknik Informatika",
	}

	fmt.Println(students)
	fmt.Println(students["nama"])
	fmt.Println(students["npm"])
	fmt.Println(students["jurusan"])
}
