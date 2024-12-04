package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	fn := float64((4*n+2)*(4*n+2)) / float64((4*n+1)*(4*n+3))
	fmt.Printf("Hasil f(N) = %.10f\n", fn)

	approximation := 1.0
	for i := 0; i <= n; i++ {
		approximation *= float64((4*i+2)*(4*i+2)) / float64((4*i+1)*(4*i+3))
	}
	fmt.Printf("Nilai perkiraan akar 2 = %.10f\n", approximation)
}
