package main

import (
	"fmt"
	"strings"
)

func kalkulasiNilai(waktu [8]int) (int, int) {
	jumlahSoal := 0
	totalNilai := 0
	for idx := 0; idx < 8; idx++ {
		if waktu[idx] <= 300 {
			jumlahSoal++
			totalNilai += waktu[idx]
		}
	}
	return jumlahSoal, totalNilai
}

func hitungJuara() {
	var (
		namaPeserta, juara        string
		durasi                    [8]int
		soalMaksimal, skorMinimal int
	)

	for {
		fmt.Print("Masukkan nama peserta (ketik 'stop' untuk keluar): ")
		fmt.Scanln()
		fmt.Scanln(&namaPeserta)

		if strings.ToLower(namaPeserta) == "stop" {
			break
		}

		fmt.Println("Masukkan durasi pengerjaan untuk 8 soal (dalam menit):")
		for idx := 0; idx < 8; idx++ {
			fmt.Scan(&durasi[idx])
		}

		jumlahSoal, totalNilai := kalkulasiNilai(durasi)

		if jumlahSoal > soalMaksimal || (jumlahSoal == soalMaksimal && totalNilai < skorMinimal) {
			juara = namaPeserta
			soalMaksimal = jumlahSoal
			skorMinimal = totalNilai
		}
	}

	if juara != "" {
		fmt.Printf("\nPemenang: %s\n", juara)
		fmt.Printf("Jumlah soal selesai: %d\n", soalMaksimal)
		fmt.Printf("Total waktu (nilai): %d\n", skorMinimal)
	} else {
		fmt.Println("Tidak ada peserta yang valid.")
	}
}

func main() {
	hitungJuara()
}
