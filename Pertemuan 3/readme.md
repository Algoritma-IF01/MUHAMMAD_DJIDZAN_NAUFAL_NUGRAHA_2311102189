# <h1 align="center"> Laporan Praktikum_3 - Modul 2 - Struktur Kontrol Perulangan & Percabangan </h1>

<p align="center"> Muhammad Djidzan Naufal Nugraha </p>

<p align="center"> 2311102189 </p>

## Guided

### Guided 1

**Source Code :**
```GO
package main

import "fmt"

func main() {
	urutanBenar := []string{"merah", "kuning", "hijau", "ungu"}
	hasil := true

	for i := 1; i <= 5; i++ {
		var warna1, warna2, warna3, warna4 string
		fmt.Printf("Percobaan %d\n", i)
		fmt.Print("Masukan warna pertama : ")
		fmt.Scanln(&warna1)
		fmt.Print("Masukan warna kedua : ")
		fmt.Scanln(&warna2)
		fmt.Print("Masukan warna ketiga : ")
		fmt.Scanln(&warna3)
		fmt.Print("Masukan warna keempat : ")
		fmt.Scanln(&warna4)

		if warna1 != urutanBenar[0] || warna2 != urutanBenar[1] || warna3 != urutanBenar[2] || warna4 != urutanBenar[3] {
			hasil = false
		}
	}
	fmt.Println("Berhasil : ", hasil)
}
```

**Screenshot Output :**
![Guided1](https://github.com/user-attachments/assets/0dd20196-98b8-4d57-b183-bfffee2d98be)

### Guided 2

**Source Code :**
```GO
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pita string
	var bungaCount int

	for {
		fmt.Printf("Bunga %d: ", bungaCount+1)
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "selesai" {
			break
		}

		if pita == "" {
			pita = input
		} else {
			pita += " – " + input
		}
		bungaCount++
	}

	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", bungaCount)
}
```

**Screenshot Output :**
![Guided2](https://github.com/user-attachments/assets/52a75cf3-160e-48c2-9819-bd3a9624ca75)

## Unguided

### Unguided 1

**Source Code :**
```GO
package main

import "fmt"

func main() {
	var totalBerat int
	fmt.Print("Masukkan berat parsel dalam gram: ")
	fmt.Scanln(&totalBerat)

	kg := totalBerat / 1000
	gram := totalBerat % 1000

	biayaPerKg := 10000
	biayaUntukKg := kg * biayaPerKg
	var biayaTambahan int

	if kg > 10 {
		biayaTambahan = 0
	} else if gram >= 500 {
		biayaTambahan = gram * 5
	} else {
		biayaTambahan = gram * 15
	}

	total := biayaUntukKg + biayaTambahan

	fmt.Printf("Rincian berat: %d kg dan %d gram\n", kg, gram)
	fmt.Printf("Biaya: Rp %d untuk kg + Rp %d untuk gram\n", biayaUntukKg, biayaTambahan)
	fmt.Printf("Total yang harus dibayar: Rp %d\n", total)
}
```

**Screenshot Output :**
![Soal1](https://github.com/user-attachments/assets/0429f747-0c7b-440d-9115-f77de8fd7254)

### Unguided 2

**Source Code :**
```GO
package main

import "fmt"

func main() {
	var nilaiAkhir float64
	var grade string

	fmt.Print("Masukkan Nilai Akhir: ")
	fmt.Scanln(&nilaiAkhir)

	switch {
	case nilaiAkhir > 80:
		grade = "A"
	case nilaiAkhir > 72.5:
		grade = "AB"
	case nilaiAkhir > 65:
		grade = "B"
	case nilaiAkhir > 57.5:
		grade = "BC"
	case nilaiAkhir > 50:
		grade = "C"
	case nilaiAkhir > 40:
		grade = "D"
	default:
		grade = "E"
	}

	fmt.Printf("Grade mata kuliah Anda: %s\n", grade)
}
```

**Screenshot Output :**
![Soal2](https://github.com/user-attachments/assets/f2c0aee5-576c-4c4b-87ec-6397e769dd9a)

### Unguided 3

**Source Code :**
```GO
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
```

**Screenshot Output :**
![Soal3](https://github.com/user-attachments/assets/c918e1db-3584-42a0-9c1a-bb4edb5a2eda)

### Unguided 4

**Source Code :**
```GO
package main

import "fmt"

func main() {
	var kantong1, kantong2 float64

	for {
		fmt.Print("Masukkan berat barang di kedua kantong (atau masukkan -1 untuk berhenti): ")
		fmt.Scan(&kantong1)

		if kantong1 == -1 {
			fmt.Println("Program dihentikan oleh pengguna.")
			break
		}

		fmt.Print("Masukkan berat kantong kedua: ")
		fmt.Scan(&kantong2)

		if kantong2 == -1 {
			fmt.Println("Program dihentikan oleh pengguna.")
			break
		}

		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Berat tidak valid. Silakan coba lagi.")
			continue
		}

		total := kantong1 + kantong2
		if total > 150 {
			fmt.Println("Berat total melebihi batas: true")
		}

		perbedaan := kantong1 - kantong2
		if perbedaan < 0 {
			perbedaan = -perbedaan
		}

		oleng := perbedaan >= 9
		fmt.Printf("Apakah motor Pak Andi akan oleng? %t\n", oleng)
	}

	fmt.Println("Proses selesai.")
}
```

**Screenshot Output :**
![Soal4](https://github.com/user-attachments/assets/41536ebd-5ead-4080-ae88-a7dd357476d1)

### Unguided 5
**Source Code :**
```GO
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
```

**Screenshot Output :**
![Soal5](https://github.com/user-attachments/assets/cf94e309-45b1-4855-ba13-67020a452032)
