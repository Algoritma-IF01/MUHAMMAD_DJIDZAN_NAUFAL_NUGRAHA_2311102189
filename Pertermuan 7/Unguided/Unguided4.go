package main

import "fmt"

type dataRumah [][]int

func urutkanGanjil(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		posisi := i
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[j] < (*arr)[posisi] {
				posisi = j
			}
		}
		(*arr)[i], (*arr)[posisi] = (*arr)[posisi], (*arr)[i]
	}
}

func urutkanGenap(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		posisi := i
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[j] > (*arr)[posisi] {
				posisi = j
			}
		}
		(*arr)[i], (*arr)[posisi] = (*arr)[posisi], (*arr)[i]
	}
}

func prosesPengurutan(rumah *dataRumah) {
	fmt.Println("\nHasil data setelah diurutkan:\n")
	for index, daftarRumah := range *rumah {
		var ganjil, genap []int

		for _, nomor := range daftarRumah {
			if nomor%2 == 0 {
				genap = append(genap, nomor)
			} else {
				ganjil = append(ganjil, nomor)
			}
		}

		urutkanGanjil(&ganjil)
		urutkanGenap(&genap)

		fmt.Printf("Daerah %d:\n", index+1)
		fmt.Print("  Ganjil: ")
		if len(ganjil) > 0 {
			for i, nomor := range ganjil {
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Print(nomor)
			}
		} else {
			fmt.Print("-")
		}
		fmt.Println()

		fmt.Print("  Genap : ")
		if len(genap) > 0 {
			for i, nomor := range genap {
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Print(nomor)
			}
		} else {
			fmt.Print("-")
		}
		fmt.Println("\n")
	}
}

func main() {
	var (
		rumahKerabat dataRumah
		totalDaerah  int
	)

	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scanln(&totalDaerah)
	fmt.Println()

	for i := 0; i < totalDaerah; i++ {
		var jumlahRumah int
		fmt.Printf("Masukkan jumlah rumah di daerah %d: ", i+1)
		fmt.Scanln(&jumlahRumah)

		var rumah []int
		fmt.Printf("Masukkan nomor rumah untuk daerah %d:\n", i+1)
		for j := 0; j < jumlahRumah; j++ {
			var nomor int
			fmt.Printf("  Rumah ke-%d: ", j+1)
			fmt.Scanln(&nomor)
			rumah = append(rumah, nomor)
		}
		fmt.Println()
		rumahKerabat = append(rumahKerabat, rumah)
	}

	fmt.Println("Data sebelum diurutkan:\n")
	for i, daerah := range rumahKerabat {
		fmt.Printf("Daerah %d: %v\n\n", i+1, daerah)
	}

	prosesPengurutan(&rumahKerabat)
}
