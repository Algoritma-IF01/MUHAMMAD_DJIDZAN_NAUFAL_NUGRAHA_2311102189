package main

import "fmt"

func main() {
	var totalBerat int
	fmt.Print("Masukkan berat parsel dalam gram: ")
	fmt.Scanln(&totalBerat)

	kg := totalBerat / 1000
	gram := totalBerat % 1000

	biayaPerKg := 10000
	biayaUntukKg := kg * biayaPerKg
	var biayaTambahan int

	if kg > 10 {
		biayaTambahan = 0
	} else if gram >= 500 {
		biayaTambahan = gram * 5
	} else {
		biayaTambahan = gram * 15
	}

	total := biayaUntukKg + biayaTambahan

	fmt.Printf("Rincian berat: %d kg dan %d gram\n", kg, gram)
	fmt.Printf("Biaya: Rp %d untuk kg + Rp %d untuk gram\n", biayaUntukKg, biayaTambahan)
	fmt.Printf("Total yang harus dibayar: Rp %d\n", total)
}
