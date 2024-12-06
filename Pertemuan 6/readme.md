<h1 align="center">Laporan Praktikum_6 - Modul 10 - Struck dan Array </h1>

<p align="center"> Muhammad Djidzan Naufal Nugraha </p>

<p align="center"> 2311102189 </p>

## Guided

### Guided 1

**Source Code :**

```GO
package main

import "fmt"

type arrInt [2023]int

// Fungsi untuk mencari indeks dari nilai terkecil
func terkecil_2(tabInt arrInt, n int) int {
	var idx int = 0 // indeks data pertama
	var j int = 1   // pencarian dimulai dari data kedua
	for j < n {
		if tabInt[idx] > tabInt[j] { // cek apakah tabInt[j] lebih kecil dari tabInt[idx]
			idx = j // update idx ke indeks baru yang nilainya lebih kecil
		}
		j = j + 1
	}
	return idx // mengembalikan indeks dari nilai terkecil
}

func main() {
	var n int
	var data arrInt

	// Input jumlah elemen N
	fmt.Print("Masukkan jumlah elemen (N <= 2023): ")
	fmt.Scan(&n)

	// Validasi N agar tidak melebihi kapasitas array
	if n <= 0 || n > 2023 {
		fmt.Println("Jumlah elemen harus antara 1 dan 2023")
		return
	}

	// Input elemen-elemen array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	// Panggil fungsi untuk mencari indeks nilai terkecil
	idxTerkecil := terkecil_2(data, n)
	fmt.Printf("Indeks nilai terkecil: %d\n", idxTerkecil)
	fmt.Printf("Nilai terkecil: %d\n", data[idxTerkecil])
}
```

**Screenshot Output :**
![Guided1](https://github.com/user-attachments/assets/54b12363-3479-428e-8ce4-71e8042b68e8)

### Guided 2

**Source Code :**

```GO
package main

import "fmt"

// Mendefinisikan tipe data mahasiswa
type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

// Mendefinisikan array mahasiswa dengan kapasitas 2023
type arrMhs [2023]mahasiswa

// Fungsi untuk mencari indeks mahasiswa dengan IPK tertinggi
func IPK_2(T arrMhs, n int) int {
	// idx menyimpan indeks mahasiswa dengan IPK tertinggi sementara
	var idx int = 0
	var j int = 1
	for j < n {
		if T[idx].ipk < T[j].ipk {
			idx = j
		}
		j = j + 1
	}
	return idx
}

func main() {
	var n int
	var data arrMhs

	// Input jumlah mahasiswa
	fmt.Print("Masukkan jumlah mahasiswa (N <= 2023): ")
	fmt.Scan(&n)

	// Validasi jumlah mahasiswa
	if n <= 0 || n > 2023 {
		fmt.Println("Jumlah mahasiswa harus antara 1 dan 2023")
		return
	}

	// Input data mahasiswa
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data mahasiswa ke-%d\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scan(&data[i].nama)
		fmt.Print("NIM: ")
		fmt.Scan(&data[i].nim)
		fmt.Print("Kelas: ")
		fmt.Scan(&data[i].kelas)
		fmt.Print("Jurusan: ")
		fmt.Scan(&data[i].jurusan)
		fmt.Print("IPK: ")
		fmt.Scan(&data[i].ipk)
	}

	// Panggil fungsi untuk mencari indeks mahasiswa dengan IPK tertinggi
	idxTertinggi := IPK_2(data, n)

	// Tampilkan data mahasiswa dengan IPK tertinggi
	fmt.Println("\nMahasiswa dengan IPK tertinggi:")
	fmt.Printf("Nama    : %s\n", data[idxTertinggi].nama)
	fmt.Printf("NIM     : %s\n", data[idxTertinggi].nim)
	fmt.Printf("Kelas   : %s\n", data[idxTertinggi].kelas)
	fmt.Printf("Jurusan : %s\n", data[idxTertinggi].jurusan)
	fmt.Printf("IPK     : %.2f\n", data[idxTertinggi].ipk)
}
```

**Screenshot Output :**
![Guided2](https://github.com/user-attachments/assets/35ecb84f-d8d7-4e64-98ab-b5f47aaaed1c)

## Unguided 

### Unguided 1

**Source Code :**

```GO
package main

import "fmt"

func main() {
	cariBeratKelinci()
}

func cariBeratKelinci() {
	var jumlah int
	fmt.Print("Masukkan jumlah kelinci yang ingin diperiksa: ")
	fmt.Scan(&jumlah)

	if jumlah < 1 || jumlah > 1000 {
		fmt.Println("Jumlah kelinci harus berada dalam rentang 1 hingga 1000.")
		return
	}

	beratKelinci := make([]float64, jumlah)
	fmt.Println("Masukkan berat masing-masing kelinci (dalam kg):")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	beratTerkecil := beratKelinci[0]
	beratTerbesar := beratKelinci[0]
	for i := 1; i < jumlah; i++ {
		if beratKelinci[i] < beratTerkecil {
			beratTerkecil = beratKelinci[i]
		}
		if beratKelinci[i] > beratTerbesar {
			beratTerbesar = beratKelinci[i]
		}
	}

	fmt.Printf("\nBerat kelinci paling ringan: %.2f kg\n", beratTerkecil)
	fmt.Printf("Berat kelinci paling berat: %.2f kg\n", beratTerbesar)
}
```

**Screenshot Output :**
![Soal1](https://github.com/user-attachments/assets/8757432c-8b85-4e7c-8cb2-70e416f4bb39)

### Unguided 2

**Source Code :**

```GO
package main

import "fmt"

func main() {
	const maxIkan = 1000

	var (
		jumlahIkan, kapasitasWadah int
		beratIkan                  [maxIkan]float64
		beratPerWadah              []float64
	)

	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah (pisahkan dengan spasi): ")
	fmt.Scanln(&jumlahIkan, &kapasitasWadah)

	fmt.Printf("\nMasukkan berat masing-masing ikan (jumlah = %d):\n", jumlahIkan)
	for i := 0; i < jumlahIkan; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	fmt.Println()
	for i := 0; i < jumlahIkan; i += kapasitasWadah {
		totalBerat := 0.0

		for j := i; j < i+kapasitasWadah && j < jumlahIkan; j++ {
			totalBerat += beratIkan[j]
		}

		beratPerWadah = append(beratPerWadah, totalBerat)
	}

	fmt.Println("Hasil penghitungan berat tiap wadah:\n")
	for i, totalBerat := range beratPerWadah {
		fmt.Printf("Wadah %d - Total berat: %.2f kg\n", i+1, totalBerat)
	}

	fmt.Println("\nRata-rata berat ikan di setiap wadah:\n")
	for i, totalBerat := range beratPerWadah {
		rataRata := totalBerat / float64(kapasitasWadah)
		fmt.Printf("Wadah %d - Rata-rata berat: %.2f kg\n", i+1, rataRata)
	}

	fmt.Println()
}
```

**Screenshot Output :**
![Soal2](https://github.com/user-attachments/assets/fa6054bc-c29f-4821-bb35-b9637cb8c799)

### Unguided 3

**Source Code :**

```GO
package main

import "fmt"

type BeratBalita [100]float64

func cariBeratTerendahTertinggi(berat BeratBalita, jumlah int, beratTerendah, beratTertinggi *float64) {
	*beratTerendah = berat[0]
	*beratTertinggi = berat[0]

	for i := 1; i < jumlah; i++ {
		if berat[i] < *beratTerendah {
			*beratTerendah = berat[i]
		}
		if berat[i] > *beratTertinggi {
			*beratTertinggi = berat[i]
		}
	}
}

func hitungRataRata(berat BeratBalita, jumlah int) float64 {
	var total float64
	for i := 0; i < jumlah; i++ {
		total += berat[i]
	}
	return total / float64(jumlah)
}

func analisisBeratBalita() {
	var jumlahBalita int
	var beratBalita BeratBalita
	var beratTerendah, beratTertinggi float64

	for {
		fmt.Print("Masukkan jumlah data berat balita: ")
		fmt.Scan(&jumlahBalita)
		if jumlahBalita > 0 && jumlahBalita <= 100 {
			break
		}
		fmt.Println("Jumlah data harus antara 1 hingga 100. Silakan coba lagi.\n")
	}

	fmt.Println()
	for i := 0; i < jumlahBalita; i++ {
		for {
			fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
			fmt.Scan(&beratBalita[i])
			if beratBalita[i] > 0 {
				break
			}
			fmt.Println("Berat balita harus positif. Silakan masukkan ulang.\n")
		}
	}

	fmt.Println()
	cariBeratTerendahTertinggi(beratBalita, jumlahBalita, &beratTerendah, &beratTertinggi)
	rataRata := hitungRataRata(beratBalita, jumlahBalita)

	fmt.Println("Hasil Analisis Berat Balita:\n")
	fmt.Printf("Berat balita terendah: %.2f kg\n", beratTerendah)
	fmt.Printf("Berat balita tertinggi: %.2f kg\n", beratTertinggi)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
	fmt.Println()
}

func main() {
	analisisBeratBalita()
}
```

**Screenshot Output :**
![Soal3](https://github.com/user-attachments/assets/ef2cb46c-4040-4dfb-a4fb-13617be4a246)

