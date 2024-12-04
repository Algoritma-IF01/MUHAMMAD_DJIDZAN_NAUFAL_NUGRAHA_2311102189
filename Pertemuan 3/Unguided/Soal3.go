package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&angka)

	if angka < 2 {
		fmt.Println("Angka harus lebih dari atau sama dengan 2")
		return
	}

	fmt.Print("Faktor dari angka tersebut: ")
	jumlahFaktor := 0

	for pembagi := 1; pembagi <= angka; pembagi++ {
		if angka%pembagi == 0 {
			fmt.Printf("%d ", pembagi)
			jumlahFaktor++
		}
	}

	isPrima := jumlahFaktor == 2
	fmt.Printf("\nApakah angka ini prima? %t\n", isPrima)
}
