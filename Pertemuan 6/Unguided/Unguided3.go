package main

import "fmt"

type BeratBalita [100]float64

func cariBeratTerendahTertinggi(berat BeratBalita, jumlah int, beratTerendah, beratTertinggi *float64) {
	*beratTerendah = berat[0]
	*beratTertinggi = berat[0]

	for i := 1; i < jumlah; i++ {
		if berat[i] < *beratTerendah {
			*beratTerendah = berat[i]
		}
		if berat[i] > *beratTertinggi {
			*beratTertinggi = berat[i]
		}
	}
}

func hitungRataRata(berat BeratBalita, jumlah int) float64 {
	var total float64
	for i := 0; i < jumlah; i++ {
		total += berat[i]
	}
	return total / float64(jumlah)
}

func analisisBeratBalita() {
	var jumlahBalita int
	var beratBalita BeratBalita
	var beratTerendah, beratTertinggi float64

	for {
		fmt.Print("Masukkan jumlah data berat balita: ")
		fmt.Scan(&jumlahBalita)
		if jumlahBalita > 0 && jumlahBalita <= 100 {
			break
		}
		fmt.Println("Jumlah data harus antara 1 hingga 100. Silakan coba lagi.\n")
	}

	fmt.Println()
	for i := 0; i < jumlahBalita; i++ {
		for {
			fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
			fmt.Scan(&beratBalita[i])
			if beratBalita[i] > 0 {
				break
			}
			fmt.Println("Berat balita harus positif. Silakan masukkan ulang.\n")
		}
	}

	fmt.Println()
	cariBeratTerendahTertinggi(beratBalita, jumlahBalita, &beratTerendah, &beratTertinggi)
	rataRata := hitungRataRata(beratBalita, jumlahBalita)

	fmt.Println("Hasil Analisis Berat Balita:\n")
	fmt.Printf("Berat balita terendah: %.2f kg\n", beratTerendah)
	fmt.Printf("Berat balita tertinggi: %.2f kg\n", beratTertinggi)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
	fmt.Println()
}

func main() {
	analisisBeratBalita()
}
