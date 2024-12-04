package main

import (
	"fmt"
	"math"
)

func hitungJarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

func dalamLingkaran(cx, cy, r, x, y int) bool {
	return hitungJarak(cx, cy, x, y) <= float64(r)
}

func cekLingkaran() {
	var (
		cx1, cy1, r1 int
		cx2, cy2, r2 int
		x, y         int
	)

	fmt.Print("Masukkan pusat dan radius lingkaran 1 (cx cy r): ")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Masukkan pusat dan radius lingkaran 2 (cx cy r): ")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Masukkan titik (x, y): ")
	fmt.Scan(&x, &y)

	dalamLingkaran1 := dalamLingkaran(cx1, cy1, r1, x, y)
	dalamLingkaran2 := dalamLingkaran(cx2, cy2, r2, x, y)

	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik berada dalam kedua lingkaran")
	} else if dalamLingkaran1 {
		fmt.Println("Titik berada dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik berada dalam lingkaran 2")
	} else {
		fmt.Println("Titik berada di luar kedua lingkaran")
	}
}

func main() {
	cekLingkaran()
}
