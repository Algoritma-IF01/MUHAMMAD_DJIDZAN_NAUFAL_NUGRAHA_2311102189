package main

import "fmt"

func kuadrat(x int) int {
	return x * x
}

func kurangDua(x int) int {
	return x - 2
}

func tambahSatu(x int) int {
	return x + 1
}

func fungsi1(n int) int {
	return kuadrat(kurangDua(tambahSatu(n)))
}

func fungsi2(m int) int {
	return kurangDua(tambahSatu(kuadrat(m)))
}

func fungsi3(p int) int {
	return tambahSatu(kuadrat(kurangDua(p)))
}

func soalKombinasi() {
	var n, m, p int
	fmt.Print("Masukkan 3  nilai : ")
	fmt.Scan(&n, &m, &p)

	fmt.Printf("\n(Nilai 1)(%d) = %d\n", n, fungsi1(n))
	fmt.Printf("(Nilai 2)(%d) = %d\n", m, fungsi2(m))
	fmt.Printf("(Nilai 3)(%d) = %d\n", p, fungsi3(p))
}

func main() {
	soalKombinasi()
}
