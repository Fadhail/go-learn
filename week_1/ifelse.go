package main

import "fmt"

func main() {
	// User Input
	fmt.Print("Masukkan nama: ")
	name := ""
	fmt.Scanln(&name)

	fmt.Print("Masukkan umur: ")
	umur := 0
	fmt.Scanln(&umur)

	if umur >= 20 && umur <= 50 {
		fmt.Println(fmt.Sprintf("Hai %s umur kamu %d, kamu diperbolehkan masuk", name, umur))
	} else if umur < 20 {
		fmt.Println(fmt.Sprintf("Hai %s umur kamu %d, kamu belum boleh masuk", name, umur))
	} else {
		fmt.Println(fmt.Sprintf("Hai %s umur kamu %d, kamu sudah tua", name, umur))
	}
}
