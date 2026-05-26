package main

import "fmt"

// Fungsi rekursif untuk mencari faktor
// Parameter i = bilangan yang sedang diperiksa, n = angka target
func cariFaktor(i, n int) {
	// Kondisi berhenti: jika bilangan pemeriksa sudah melebihi angka target
	if i > n {
		return // Keluar dari fungsi, tidak melakukan apa-apa lagi
	}
	// Pengecekan: jika n habis dibagi i, berarti i adalah faktor dari n
	if n%i == 0 {
		fmt.Printf("%d ", i) // Tampilkan angka faktor tersebut
	}
	// Panggil fungsi lagi dengan nilai i ditambah 1 (periksa angka berikutnya)
	cariFaktor(i+1, n)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n) // Membaca angka yang akan dicari faktornya

	fmt.Print("Keluaran: ")
	cariFaktor(1, n) // Memulai pencarian dari angka 1 sampai n
	fmt.Println()
}
