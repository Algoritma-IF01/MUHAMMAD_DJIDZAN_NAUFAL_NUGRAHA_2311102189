package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const maksimalElemen int = 127

type ArrayKarakter [maksimalElemen]rune

func masukkanTeks(array *ArrayKarakter, panjang *int) bool {
	fmt.Print("Masukkan teks (akhiri dengan . atau ketik STOP untuk berhenti):\n")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	teks := scanner.Text()

	if strings.ToUpper(teks) == "STOP" {
		return true
	}

	for _, karakter := range teks {
		if karakter == '.' || *panjang >= maksimalElemen {
			break
		}

		array[*panjang] = karakter
		(*panjang)++
	}
	fmt.Println()
	return false
}

func tampilkanArray(array ArrayKarakter, panjang int) {
	for i := 0; i < panjang; i++ {
		fmt.Printf("%c", array[i])
	}
	fmt.Println("\n")
}

func balikkanArray(array *ArrayKarakter, panjang int) {
	for i := 0; i < panjang/2; i++ {
		array[i], array[panjang-1-i] = array[panjang-1-i], array[i]
	}
}

func cekPalindrom(array ArrayKarakter, panjang int) bool {
	for i := 0; i < panjang/2; i++ {
		if array[i] != array[panjang-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var teks ArrayKarakter
	var panjangTeks int

	for {
		berhenti := masukkanTeks(&teks, &panjangTeks)
		if berhenti {
			fmt.Println("Program dihentikan oleh pengguna.")
			break
		}

		balikkanArray(&teks, panjangTeks)
		fmt.Print("Teks setelah dibalik:\n")
		tampilkanArray(teks, panjangTeks)

		apakahPalindrom := cekPalindrom(teks, panjangTeks)
		fmt.Printf("Apakah teks merupakan palindrom?\n%t\n", apakahPalindrom)
	}
}
