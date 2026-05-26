package main

import (
	"fmt"
)

func main() {
	// Deklarasi array untuk menampung berat kelinci, maksimal 1000 data
	var beratKelinci [1000]float64
	var n int

	// Membaca jumlah data yang akan dimasukkan
	fmt.Scan(&n)

	// Membatasi jumlah data agar tidak melebihi kapasitas array
	if n > 1000 {
		n = 1000
	}

	// Membaca data berat kelinci sebanyak n kali
	for i := 0; i < n; i++ {
		fmt.Scan(&beratKelinci[i])
	}

	// Inisialisasi nilai awal min dan max dengan data pertama
	min := beratKelinci[0]
	max := beratKelinci[0]

	// Pencarian nilai terkecil dan terbesar
	for i := 1; i < n; i++ {
		if beratKelinci[i] < min {
			min = beratKelinci[i]
		}
		if beratKelinci[i] > max {
			max = beratKelinci[i]
		}
	}

	// Menampilkan hasil dengan format 2 angka di belakang koma
	fmt.Printf("%.2f %.2f\n", min, max)
}
