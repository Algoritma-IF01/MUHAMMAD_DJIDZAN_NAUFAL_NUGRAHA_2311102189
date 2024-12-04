package main

import "fmt"

func main() {
	var nilaiAkhir float64
	var grade string

	fmt.Print("Masukkan Nilai Akhir: ")
	fmt.Scanln(&nilaiAkhir)

	switch {
	case nilaiAkhir > 80:
		grade = "A"
	case nilaiAkhir > 72.5:
		grade = "AB"
	case nilaiAkhir > 65:
		grade = "B"
	case nilaiAkhir > 57.5:
		grade = "BC"
	case nilaiAkhir > 50:
		grade = "C"
	case nilaiAkhir > 40:
		grade = "D"
	default:
		grade = "E"
	}

	fmt.Printf("Grade mata kuliah Anda: %s\n", grade)
}
