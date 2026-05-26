package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n) // Membaca angka batas akhir dari pengguna
	// Memanggil fungsi penjumlahan dan langsung mencetak hasilnya
	fmt.Println(penjumlahan(n))
}

// Fungsi rekursif untuk menjumlahkan angka dari 1 hingga n
func penjumlahan(n int) int {
	// Kondisi berhenti: jika n bernilai 1, kembalikan nilai 1
	if n == 1 {
		return 1
	} else {
		// Rumus: angka saat ini + hasil penjumlahan angka sebelumnya
		return n + penjumlahan(n-1)
	}
}
