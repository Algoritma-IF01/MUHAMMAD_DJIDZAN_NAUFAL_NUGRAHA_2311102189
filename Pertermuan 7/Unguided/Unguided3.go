package main

import (
	"fmt"
)

const kapasitasMaksimal = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	salinan, tahun, rating       int
}

type KoleksiBuku [kapasitasMaksimal]Buku

func TambahkanBuku(koleksi *KoleksiBuku, jumlah *int) {
	fmt.Print("Masukkan jumlah buku yang akan didaftarkan: ")
	fmt.Scanln(jumlah)

	for i := 0; i < *jumlah; i++ {
		fmt.Printf("\nMasukkan data untuk buku ke-%d\n", i+1)

		for {
			fmt.Print("Masukkan ID buku: ")
			fmt.Scanln(&(*koleksi)[i].id)

			duplikat := false
			for j := 0; j < i; j++ {
				if (*koleksi)[j].id == (*koleksi)[i].id {
					duplikat = true
					break
				}
			}

			if duplikat {
				fmt.Println("ID buku sudah digunakan, silakan masukkan ID lain.")
			} else {
				break
			}
		}

		fmt.Print("Masukkan judul buku: ")
		fmt.Scanln(&(*koleksi)[i].judul)

		fmt.Print("Masukkan penulis buku: ")
		fmt.Scanln(&(*koleksi)[i].penulis)

		fmt.Print("Masukkan penerbit buku: ")
		fmt.Scanln(&(*koleksi)[i].penerbit)

		fmt.Print("Masukkan jumlah salinan buku: ")
		fmt.Scanln(&(*koleksi)[i].salinan)

		fmt.Print("Masukkan tahun penerbitan buku: ")
		fmt.Scanln(&(*koleksi)[i].tahun)

		for {
			fmt.Print("Masukkan rating buku (1-10): ")
			fmt.Scanln(&(*koleksi)[i].rating)
			if (*koleksi)[i].rating < 1 || (*koleksi)[i].rating > 10 {
				fmt.Println("Rating harus berada di antara 1 hingga 10.")
			} else {
				break
			}
		}
	}
	fmt.Println()
}

func TampilkanBukuTerbaik(koleksi KoleksiBuku, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada buku yang didaftarkan.\n")
		return
	}

	maxRating := koleksi[0].rating
	idx := 0
	for i := 1; i < jumlah; i++ {
		if koleksi[i].rating > maxRating {
			maxRating = koleksi[i].rating
			idx = i
		}
	}

	fmt.Println("\nBuku dengan Rating Tertinggi:")
	fmt.Printf("Judul     : %s\n", koleksi[idx].judul)
	fmt.Printf("Penulis   : %s\n", koleksi[idx].penulis)
	fmt.Printf("Penerbit  : %s\n", koleksi[idx].penerbit)
	fmt.Printf("Tahun     : %d\n", koleksi[idx].tahun)
	fmt.Printf("Salinan   : %d\n", koleksi[idx].salinan)
	fmt.Printf("Rating    : %d\n\n", koleksi[idx].rating)
}

func UrutkanBukuBerdasarkanRating(koleksi *KoleksiBuku, jumlah int) {
	for i := 1; i < jumlah; i++ {
		temp := (*koleksi)[i]
		j := i - 1
		for j >= 0 && (*koleksi)[j].rating < temp.rating {
			(*koleksi)[j+1] = (*koleksi)[j]
			j--
		}
		(*koleksi)[j+1] = temp
	}
}

func Tampilkan5BukuTeratas(koleksi KoleksiBuku, jumlah int) {
	fmt.Println("\n5 Buku dengan Rating Tertinggi:")
	maks := 5
	if jumlah < 5 {
		maks = jumlah
	}
	for i := 0; i < maks; i++ {
		fmt.Printf("%d. %s\n", i+1, koleksi[i].judul)
	}
	fmt.Println()
}

func CariBukuDenganRating(koleksi KoleksiBuku, jumlah int, rating int) {
	low, high := 0, jumlah-1
	for low <= high {
		mid := (low + high) / 2
		if koleksi[mid].rating == rating {
			fmt.Println("\nBuku Ditemukan:")
			fmt.Printf("Judul     : %s\n", koleksi[mid].judul)
			fmt.Printf("Penulis   : %s\n", koleksi[mid].penulis)
			fmt.Printf("Penerbit  : %s\n", koleksi[mid].penerbit)
			fmt.Printf("Tahun     : %d\n", koleksi[mid].tahun)
			fmt.Printf("Salinan   : %d\n", koleksi[mid].salinan)
			fmt.Printf("Rating    : %d\n\n", koleksi[mid].rating)
			return
		} else if koleksi[mid].rating < rating {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating tersebut.\n")
}

func main() {
	var koleksi KoleksiBuku
	var jumlahBuku, rating int

	TambahkanBuku(&koleksi, &jumlahBuku)
	TampilkanBukuTerbaik(koleksi, jumlahBuku)
	UrutkanBukuBerdasarkanRating(&koleksi, jumlahBuku)
	Tampilkan5BukuTeratas(koleksi, jumlahBuku)
	fmt.Print("Masukkan rating yang ingin dicari: ")
	fmt.Scanln(&rating)
	CariBukuDenganRating(koleksi, jumlahBuku, rating)
}
