package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	// Iterasi untuk menyusun array
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		// Pindahkan elemen yang lebih besar dari key ke kanan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func cekJarakKonsisten(arr []int) bool {
	if len(arr) < 2 {
		return true // Data dengan kurang dari 2 elemen selalu konsisten
	}
	selisih := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != selisih {
			return false
		}
	}
	return true
}

func main() {
	var input []int
	var nilai int

	fmt.Println("Masukkan angka-angka (masukkan angka negatif untuk berhenti):")
	for fmt.Scan(&nilai); nilai >= 0; fmt.Scan(&nilai) {
		input = append(input, nilai)
	}

	if len(input) == 0 {
		fmt.Println("Tidak ada data yang dimasukkan.\n")
		return
	}

	insertionSort(input)

	fmt.Println("\nHasil setelah disortir:")
	for _, angka := range input {
		fmt.Printf("%d ", angka)
	}
	fmt.Println("\n")

	if cekJarakKonsisten(input) {
		fmt.Printf("Status: Data memiliki jarak tetap %d.\n\n", input[1]-input[0])
	} else {
		fmt.Println("Status: Data tidak memiliki jarak tetap.\n")
	}
}
