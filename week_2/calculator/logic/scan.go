package logic

import "fmt"

func GetInput() (int, int, string) {
	fmt.Print("Masukkan angka pertama: ")
	var angka1 int
	fmt.Scanln(&angka1)

	fmt.Print("Masukkan angka kedua: ")
	var angka2 int
	fmt.Scanln(&angka2)

	fmt.Print("Masukkan Operator (+, -, *, /): ")
	var operasi string
	fmt.Scanln(&operasi)

	return angka1, angka2, operasi
}
