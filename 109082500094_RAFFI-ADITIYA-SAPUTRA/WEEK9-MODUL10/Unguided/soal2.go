package main

import (
	"fmt"
)

func main() {
	// Array menyimpan berat ikan, kapasitas maksimal 1000
	var beratIkan [1000]float64
	var x, y int

	// Input: x = jumlah ikan, y = kapasitas per wadah
	fmt.Scan(&x, &y)

	// Membatasi jumlah ikan agar tidak melebihi kapasitas array
	if x > 1000 {
		x = 1000
	}

	// Membaca berat masing-masing ikan
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	// Array menyimpan total berat di setiap wadah
	var totalBeratWadah [1000]float64
	var jumlahWadah int = 0

	// Proses memasukkan ikan ke wadah dan menjumlahkan beratnya
	for i := 0; i < x; i++ {
		// Menentukan nomor wadah: ikan ke-0..y-1 -> wadah 0, ikan ke-y..2y-1 -> wadah 1, dst.
		indeksWadah := i / y
		totalBeratWadah[indeksWadah] += beratIkan[i]

		// Menghitung berapa banyak wadah yang terpakai
		if indeksWadah >= jumlahWadah {
			jumlahWadah = indeksWadah + 1
		}
	}

	// Menampilkan TOTAL BERAT setiap wadah
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("%.2f", totalBeratWadah[i])
		if i < jumlahWadah-1 {
			fmt.Print(" ") // Beri spasi antar angka, tidak di akhir
		}
	}
	fmt.Println() // Pindah baris

	// Menampilkan RATA-RATA BERAT setiap wadah
	for i := 0; i < jumlahWadah; i++ {
		var ikanDiWadahIni int

		// Cek wadah terakhir jika isinya kurang dari kapasitas y
		if i == jumlahWadah-1 && x%y != 0 {
			ikanDiWadahIni = x % y
		} else {
			ikanDiWadahIni = y
		}

		// Hitung rata-rata
		rataRata := totalBeratWadah[i] / float64(ikanDiWadahIni)
		fmt.Printf("%.2f", rataRata)
		if i < jumlahWadah-1 {
			fmt.Print(" ") // Beri spasi antar angka, tidak di akhir
		}
	}
	fmt.Println()
}
