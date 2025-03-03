package curency

import "fmt"

func Converter() {
	aCountry, tCountry, money := GetValue()

	rates := Currency()

	countryToCurrency := map[string]string{
		"Indonesia": "Rupiah",
		"Amerika":   "USD",
		"Eropa":     "Euro",
	}

	fromCurrency, fromExists := countryToCurrency[aCountry]
	toCurrency, toExists := countryToCurrency[tCountry]

	if !fromExists || !toExists {
		fmt.Println("Negara tidak ditemukan dalam daftar.")
		return
	}

	fromRate, fromRateExists := rates[fromCurrency]
	toRate, toRateExists := rates[toCurrency]

	if !fromRateExists || !toRateExists {
		fmt.Println("Kurs mata uang tidak tersedia.")
		return
	}

	convertedAmount := (money / fromRate) * toRate
	fmt.Printf("Uang Anda: %.2f %s\n", convertedAmount, toCurrency)
}
