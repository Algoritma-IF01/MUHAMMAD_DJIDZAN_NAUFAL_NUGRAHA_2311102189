package main

import (
	"fmt"
	"math"
)

func analisaArray() {
	const batasMaks = 100
	var jumlahElemen, kelipatan, hapusIdx, bilangan int

	for {
		fmt.Print("Jumlah elemen array (maksimal 100): ")
		fmt.Scan(&jumlahElemen)

		if jumlahElemen > batasMaks {
			fmt.Printf("Jumlah elemen tidak boleh melebihi %d. Silakan coba lagi.\n", batasMaks)
		} else {
			break
		}
	}

	array := make([]int, jumlahElemen)

	fmt.Println("\nMasukkan nilai elemen array:")
	for idx := 0; idx < jumlahElemen; idx++ {
		fmt.Printf("Elemen ke-%d: ", idx)
		fmt.Scan(&array[idx])
	}

	fmt.Println("\nArray:")
	for _, nilai := range array {
		fmt.Printf("%d ", nilai)
	}
	fmt.Println()

	fmt.Println("\nElemen ganjil:")
	for _, nilai := range array {
		if nilai%2 != 0 {
			fmt.Printf("%d ", nilai)
		}
	}
	fmt.Println()

	fmt.Println("\nElemen genap:")
	for _, nilai := range array {
		if nilai%2 == 0 {
			fmt.Printf("%d ", nilai)
		}
	}
	fmt.Println()

	fmt.Print("\nMasukkan kelipatan indeks untuk ditampilkan: ")
	fmt.Scan(&kelipatan)
	fmt.Printf("\nElemen pada indeks kelipatan %d:\n", kelipatan)
	for idx, nilai := range array {
		if idx%kelipatan == 0 {
			fmt.Printf("%d ", nilai)
		}
	}
	fmt.Println()

	fmt.Print("\nIndeks elemen yang ingin dihapus: ")
	fmt.Scan(&hapusIdx)
	if hapusIdx >= 0 && hapusIdx < jumlahElemen {
		for i := hapusIdx; i < jumlahElemen-1; i++ {
			array[i] = array[i+1]
		}
		jumlahElemen--

		fmt.Println("\nArray setelah penghapusan:")
		for i := 0; i < jumlahElemen; i++ {
			fmt.Printf("%d ", array[i])
		}
		fmt.Println()
	} else {
		fmt.Println("\nIndeks yang dimasukkan tidak valid!")
	}

	total := 0
	for _, nilai := range array {
		total += nilai
	}
	rataRata := float64(total) / float64(jumlahElemen)
	fmt.Printf("\nRata-rata: %.2f\n", rataRata)

	var deviasiStandar float64
	for _, nilai := range array {
		deviasiStandar += math.Pow(float64(nilai)-rataRata, 2)
	}
	deviasiStandar = math.Sqrt(deviasiStandar / float64(jumlahElemen))
	fmt.Printf("Deviasi standar: %.2f\n", deviasiStandar)

	fmt.Print("\nMasukkan bilangan untuk mencari frekuensinya: ")
	fmt.Scan(&bilangan)
	frekuensi := 0
	for _, nilai := range array {
		if nilai == bilangan {
			frekuensi++
		}
	}
	fmt.Printf("\nFrekuensi bilangan %d: %d\n", bilangan, frekuensi)

	fmt.Println("\nProgram selesai.")
}

func main() {
	analisaArray()
}
