package main

import "fmt"

func main() {
	const maxIkan = 1000

	var (
		jumlahIkan, kapasitasWadah int
		beratIkan                  [maxIkan]float64
		beratPerWadah              []float64
	)

	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah (pisahkan dengan spasi): ")
	fmt.Scanln(&jumlahIkan, &kapasitasWadah)

	fmt.Printf("\nMasukkan berat masing-masing ikan (jumlah = %d):\n", jumlahIkan)
	for i := 0; i < jumlahIkan; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	fmt.Println()
	for i := 0; i < jumlahIkan; i += kapasitasWadah {
		totalBerat := 0.0

		for j := i; j < i+kapasitasWadah && j < jumlahIkan; j++ {
			totalBerat += beratIkan[j]
		}

		beratPerWadah = append(beratPerWadah, totalBerat)
	}

	fmt.Println("Hasil penghitungan berat tiap wadah:\n")
	for i, totalBerat := range beratPerWadah {
		fmt.Printf("Wadah %d - Total berat: %.2f kg\n", i+1, totalBerat)
	}

	fmt.Println("\nRata-rata berat ikan di setiap wadah:\n")
	for i, totalBerat := range beratPerWadah {
		rataRata := totalBerat / float64(kapasitasWadah)
		fmt.Printf("Wadah %d - Rata-rata berat: %.2f kg\n", i+1, rataRata)
	}

	fmt.Println()
}
