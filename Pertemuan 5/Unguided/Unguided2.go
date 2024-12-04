package main

import "fmt"

func pertandinganSepakBola() {
	var tim1, tim2 string
	var skorTim1, skorTim2 int
	var hasil []string
	nomorPertandingan := 1
	const maxPertandingan = 10

	fmt.Print("Nama Tim 1: ")
	fmt.Scanln(&tim1)
	fmt.Print("Nama Tim 2: ")
	fmt.Scanln(&tim2)
	fmt.Println()

	for nomorPertandingan <= maxPertandingan {
		fmt.Printf("Pertandingan %d (format: skor_tim1 skor_tim2): ", nomorPertandingan)
		fmt.Scanln(&skorTim1, &skorTim2)

		if skorTim1 < 0 || skorTim2 < 0 {
			fmt.Println("\nSkor negatif ditemukan, pertandingan selesai.")
			break
		}

		if skorTim1 > skorTim2 {
			hasil = append(hasil, fmt.Sprintf("Pemenang: %s", tim1))
		} else if skorTim1 < skorTim2 {
			hasil = append(hasil, fmt.Sprintf("Pemenang: %s", tim2))
		} else {
			hasil = append(hasil, "Hasil: Seri")
		}

		nomorPertandingan++
		fmt.Println()
	}

	fmt.Println("\nHasil pertandingan:")
	for i, h := range hasil {
		fmt.Printf("Pertandingan %d: %s\n", i+1, h)
	}

	fmt.Println("\nProgram selesai.")
}

func main() {
	pertandinganSepakBola()
}
