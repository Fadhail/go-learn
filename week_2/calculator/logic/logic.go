package logic

import "fmt"

func Logic() {
	angka1, angka2, operasi := GetInput()

	// Operator
	tambah := angka1 + angka2
	kurang := angka1 - angka2
	kali := angka1 * angka2
	bagi := angka1 / angka2

	if operasi == "+" {
		fmt.Println("Hasil =", tambah)
	} else if operasi == "-" {
		fmt.Println("Hasil =", kurang)
	} else if operasi == "*" {
		fmt.Println("Hasil =", kali)
	} else if operasi == "/" {
		fmt.Println("Hasil =", bagi)
	} else {
		fmt.Println("Operator Tidak Valid")
	}
}
