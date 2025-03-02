package main

import "fmt"

func main() {
	fmt.Print("Masukkan nama: ")
	name := ""
	fmt.Scanln(&name)
	var nameLength int = len(name)

	switch {
	case nameLength < 5:
		fmt.Println("Nama terlalu pendek")
	case nameLength > 5:
		fmt.Println("Nama terlau panjang")
	case nameLength == 5:
		fmt.Println("Nama pas")
	default:
		fmt.Println("Nama tidak valid")
	}
}
