package main

import "fmt"

func main() {
	cariBeratKelinci()
}

func cariBeratKelinci() {
	var jumlah int
	fmt.Print("Masukkan jumlah kelinci yang ingin diperiksa: ")
	fmt.Scan(&jumlah)

	if jumlah < 1 || jumlah > 1000 {
		fmt.Println("Jumlah kelinci harus berada dalam rentang 1 hingga 1000.")
		return
	}

	beratKelinci := make([]float64, jumlah)
	fmt.Println("Masukkan berat masing-masing kelinci (dalam kg):")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	beratTerkecil := beratKelinci[0]
	beratTerbesar := beratKelinci[0]
	for i := 1; i < jumlah; i++ {
		if beratKelinci[i] < beratTerkecil {
			beratTerkecil = beratKelinci[i]
		}
		if beratKelinci[i] > beratTerbesar {
			beratTerbesar = beratKelinci[i]
		}
	}

	fmt.Printf("\nBerat kelinci paling ringan: %.2f kg\n", beratTerkecil)
	fmt.Printf("Berat kelinci paling berat: %.2f kg\n", beratTerbesar)
}
