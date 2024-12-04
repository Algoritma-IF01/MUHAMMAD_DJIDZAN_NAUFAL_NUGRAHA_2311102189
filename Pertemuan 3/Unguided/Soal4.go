package main

import "fmt"

func main() {
	var kantong1, kantong2 float64

	for {
		fmt.Print("Masukkan berat barang di kedua kantong (atau masukkan -1 untuk berhenti): ")
		fmt.Scan(&kantong1)

		if kantong1 == -1 {
			fmt.Println("Program dihentikan oleh pengguna.")
			break
		}

		fmt.Print("Masukkan berat kantong kedua: ")
		fmt.Scan(&kantong2)

		if kantong2 == -1 {
			fmt.Println("Program dihentikan oleh pengguna.")
			break
		}

		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Berat tidak valid. Silakan coba lagi.")
			continue
		}

		total := kantong1 + kantong2
		if total > 150 {
			fmt.Println("Berat total melebihi batas: true")
		}

		perbedaan := kantong1 - kantong2
		if perbedaan < 0 {
			perbedaan = -perbedaan
		}

		oleng := perbedaan >= 9
		fmt.Printf("Apakah motor Pak Andi akan oleng? %t\n", oleng)
	}

	fmt.Println("Proses selesai.")
}
