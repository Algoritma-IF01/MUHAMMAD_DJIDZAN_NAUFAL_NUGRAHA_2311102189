package main

import (
	"fmt"
)

func selectionSort(array []int) {
	// Lakukan iterasi untuk mencari elemen terkecil dan menukarnya
	for i := 0; i < len(array)-1; i++ {
		indexMin := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[indexMin] {
				indexMin = j
			}
		}
		// Tukar elemen
		array[i], array[indexMin] = array[indexMin], array[i]
	}
}

func prosesDataDaerah() {
	var jumlahDaerah int
	fmt.Print("Masukkan jumlah daerah kerabat: ")
	fmt.Scan(&jumlahDaerah)

	if jumlahDaerah < 1 || jumlahDaerah > 999 {
		fmt.Println("Jumlah daerah kerabat harus di antara 1 dan 999.\n")
		return
	}

	dataRumah := make([][]int, jumlahDaerah)

	for i := 0; i < jumlahDaerah; i++ {
		var jumlahRumah int
		fmt.Printf("\nMasukkan jumlah rumah di daerah %d: ", i+1)
		fmt.Scan(&jumlahRumah)

		if jumlahRumah < 1 || jumlahRumah > 999999 {
			fmt.Println("Jumlah rumah di daerah harus di antara 1 dan 999999.\n")
			return
		}

		fmt.Print("Masukkan nomor-nomor rumah kerabat: ")
		rumah := make([]int, jumlahRumah)
		for j := 0; j < jumlahRumah; j++ {
			fmt.Scan(&rumah[j])
		}
		selectionSort(rumah)
		dataRumah[i] = rumah
	}

	fmt.Println("\nHasil data setelah diurutkan:")
	for idx, rumahTersortir := range dataRumah {
		fmt.Printf("Daerah %d:\n", idx+1)
		for _, nomor := range rumahTersortir {
			fmt.Printf("%d ", nomor)
		}
		fmt.Println("\n") // Tambahkan baris kosong setelah setiap daerah
	}
}

func main() {
	prosesDataDaerah()
}
