package main

import (
	"fmt"
)

func buatDeret(x int) {
	if x >= 1000000 {
		fmt.Println("Bilangan harus lebih kecil dari 1000000")
		return
	}

	fmt.Print(x, " ")

	for x != 1 {
		if x%2 == 0 {
			x = x / 2
		} else {
			x = 3*x + 1
		}
		fmt.Print(x, " ")
	}
}

func main() {
	var angka int

	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&angka)

	buatDeret(angka)
}
