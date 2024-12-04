<h1 align="center"> Laporan Praktikum_4 - Modul 3 & 4 - Fungsi dan Prosedur </h1>

<p align="center"> Muhammad Djidzan Naufal Nugraha </p>

<p align="center"> 2311102189 </p>

## Guided

### Guided 1
**Source Code :**

```GO
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println(permutasi(a, b))
	} else {
		fmt.Println(permutasi(b, a))
	}
}

func faktorial(n int) int {
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
```

**Screenshot Output :**    
![Guided1](https://github.com/user-attachments/assets/ac9a0941-58f6-4dc9-b8c3-d7b9e3fae1e4)

### Guided 2

**Source Code :**

```GO
package main

import "fmt"

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n int, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n int, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Println("Masukkan nilai a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	p1 := permutasi(a, c)
	c1 := kombinasi(a, c)

	p2 := permutasi(b, d)
	c2 := kombinasi(b, d)

	fmt.Printf("P(%d, %d) = %d\n", a, c, p1)
	fmt.Printf("C(%d, %d) = %d\n", a, c, c1)
	fmt.Printf("P(%d, %d) = %d\n", b, d, p2)
	fmt.Printf("C(%d, %d) = %d\n", b, d, c2)
}
```

**Screenshot Output :**   
![Guided2](https://github.com/user-attachments/assets/429fe9a7-102c-408e-8456-3cc96cf7e02b)

### Guided 3

**Source Code :**

```GO
package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	fmt.Println("Deret Fibonacci hingga suku ke-10: ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("Fibonacci(%d)= %d\n", i, fibonacci(i))
	}
}
```

**Screenshot Output :** 
![Guided3](https://github.com/user-attachments/assets/446f0672-2114-4dd7-b129-9a9edbcdec99)

## Unguided

### Unguided 1

**Source Code :**

```GO
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
```

**Screenshot Output :** 
![Soal1](https://github.com/user-attachments/assets/99c54373-e2c6-4eda-9a24-37a4088387c4)

### Unguided 2

**Source Code :**

```GO
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
```

**Screenshot Output :** 
![Soal2](https://github.com/user-attachments/assets/5395bc2a-06eb-44a1-b12a-7c301bec3cc1)

### Unguided 3

**Source Code :**

```GO
package main

import (
	"fmt"
	"strings"
)

func kalkulasiNilai(waktu [8]int) (int, int) {
	jumlahSoal := 0
	totalNilai := 0
	for idx := 0; idx < 8; idx++ {
		if waktu[idx] <= 300 {
			jumlahSoal++
			totalNilai += waktu[idx]
		}
	}
	return jumlahSoal, totalNilai
}

func hitungJuara() {
	var (
		namaPeserta, juara        string
		durasi                    [8]int
		soalMaksimal, skorMinimal int
	)

	for {
		fmt.Print("Masukkan nama peserta (ketik 'stop' untuk keluar): ")
		fmt.Scanln()
		fmt.Scanln(&namaPeserta)

		if strings.ToLower(namaPeserta) == "stop" {
			break
		}

		fmt.Println("Masukkan durasi pengerjaan untuk 8 soal (dalam menit):")
		for idx := 0; idx < 8; idx++ {
			fmt.Scan(&durasi[idx])
		}

		jumlahSoal, totalNilai := kalkulasiNilai(durasi)

		if jumlahSoal > soalMaksimal || (jumlahSoal == soalMaksimal && totalNilai < skorMinimal) {
			juara = namaPeserta
			soalMaksimal = jumlahSoal
			skorMinimal = totalNilai
		}
	}

	if juara != "" {
		fmt.Printf("\nPemenang: %s\n", juara)
		fmt.Printf("Jumlah soal selesai: %d\n", soalMaksimal)
		fmt.Printf("Total waktu (nilai): %d\n", skorMinimal)
	} else {
		fmt.Println("Tidak ada peserta yang valid.")
	}
}

func main() {
	hitungJuara()
}
```

**Screenshot Output :** 
![Soal3](https://github.com/user-attachments/assets/7660e72b-1ae0-44b6-96a2-bc45dbb78820)

###   Unguided 4

**Source Code :**

```GO
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
```

**Screenshot Output :** 
![Soal4](https://github.com/user-attachments/assets/791730c9-a553-45f8-b795-1131d1418aeb)
