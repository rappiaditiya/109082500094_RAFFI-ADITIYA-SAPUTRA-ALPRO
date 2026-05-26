package main

import "fmt"

// Fungsi rekursif untuk mencari dan mencetak bilangan ganjil
// Parameter i = angka yang sedang diperiksa, n = batas akhir angka
func cetakBilanganGanjil(i, n int) {
	// Kondisi berhenti: jika angka yang diperiksa sudah melebihi batas akhir
	if i > n {
		return // Menghentikan proses rekursi
	}
	// Pengecekan: jika angka i dibagi 2 sisanya bukan 0, berarti itu bilangan ganjil
	if i%2 != 0 {
		fmt.Printf("%d ", i) // Tampilkan angka tersebut
	}
	// Panggil fungsi lagi untuk memeriksa angka berikutnya
	cetakBilanganGanjil(i+1, n)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n) // Membaca batas akhir angka dari pengguna

	fmt.Print("Keluaran: ")
	cetakBilanganGanjil(1, n) // Mulai pemeriksaan dari angka 1
	fmt.Println()
}
