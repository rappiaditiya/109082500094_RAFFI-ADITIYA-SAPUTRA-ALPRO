package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n) // Membaca angka yang dimasukkan pengguna
	baris(n)     // Memanggil fungsi rekursif dengan nilai awal n
}

// Fungsi rekursif: memanggil dirinya sendiri sampai kondisi berhenti terpenuhi
func baris(bilangan int) {
	// Kondisi berhenti: jika bilangan sama dengan 1, cetak 1 dan berhenti
	if bilangan == 1 {
		fmt.Println(1)
	} else {
		// Cetak nilai saat ini
		fmt.Println(bilangan)
		// Panggil fungsi ini lagi dengan nilai yang dikurangi 1
		baris(bilangan - 1)
	}
}
