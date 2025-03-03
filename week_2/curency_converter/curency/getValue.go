package curency

import "fmt"

func GetValue() (string, string, float64) {
	fmt.Print("Pilih Negara Asal: ")
	var aCountry string
	fmt.Scanln(&aCountry)

	fmt.Print("Pilih Negara Tujuan: ")
	var tCountry string
	fmt.Scanln(&tCountry)

	fmt.Print("Nominal Uang: ")
	var money float64
	fmt.Scanln(&money)

	return aCountry, tCountry, money
}
