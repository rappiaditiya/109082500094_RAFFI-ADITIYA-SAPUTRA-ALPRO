package main

import "fmt"

func main() {
	// --- CONTOH ARRAY (Statis) ---
	// Ukuran [3] berarti jumlah datanya tetap, tidak bisa ditambah atau dikurangi
	var rakSepatu = [3]string{"Nike", "Adidas", "Jordan"}
	fmt.Println("Array:", rakSepatu)

	// --- CONTOH SLICE (Dinamis) ---
	// Dideklarasikan tanpa ukuran, datanya bisa bertambah
	var keranjangBelanja []string

	// Menambah data ke dalam slice menggunakan append
	keranjangBelanja = append(keranjangBelanja, "Apel")
	keranjangBelanja = append(keranjangBelanja, "Mangga", "Jeruk")

	fmt.Println("Slice awal:", keranjangBelanja)
	fmt.Println("Jumlah data (len):", len(keranjangBelanja)) // Menghitung jumlah isi data

	// Mengambil sebagian data dari slice (Slicing)
	// Format: [indeks_awal : indeks_akhir]
	// Ambil dari indeks 0 sampai sebelum indeks 2 (data indeks 0 dan 1)
	potongan := keranjangBelanja[0:2]
	fmt.Println("Hasil Slicing [0:2]:", potongan)
}
