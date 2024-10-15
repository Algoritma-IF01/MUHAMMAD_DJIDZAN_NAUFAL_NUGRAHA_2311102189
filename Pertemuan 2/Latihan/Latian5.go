package main

import (
	"fmt"
	// "strconv"
)

func main() {
	var num int

	fmt.Print("Masukkan 5 angka dipisahkan spasi: ")
	for i := 0; i < 5; i++ {
		fmt.Scan(&num)
		fmt.Printf("%c", num)
	}
	fmt.Println()

	var input string
	fmt.Print("Masukkan 3 karakter: ")
	fmt.Scan(&input)

	if len(input) <= 3 {
		fmt.Print("Hasil konversi: ")
		for i := 0; i < len(input); i++ {
			fmt.Printf("%c", input[i]+1)
		}
		fmt.Println()
	}
}
