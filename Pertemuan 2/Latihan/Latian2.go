package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Masukan Tahun: ")
	fmt.Scanln(&tahun)

	kabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	fmt.Printf("Kabisat: %t", kabisat)
}
