package main

import "fmt"

func main() {
	var r int
	fmt.Print("Jari-Jari: ")
	fmt.Scanln(&r)

	pi := 3.1415926535
	volume := (4.0 / 3.0) * pi * float64(r*r*r)
	luas := 4 * pi * float64(r*r)

	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f", r, volume, luas)
}