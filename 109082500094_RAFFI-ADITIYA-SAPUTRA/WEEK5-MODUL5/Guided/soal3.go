package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)            // Membaca nilai pangkat yang diinginkan dari pengguna
	fmt.Println(pangkat(n)) // Memanggil fungsi dan mencetak hasilnya
}

func pangkat(n int) int {
	// Kondisi berhenti: apa pun angka yang dipangkatkan 0 hasilnya pasti 1
	if n == 0 {
		return 1
	} else {
		// Rumus: 2 dikali dengan hasil 2 pangkat sebelumnya
		return 2 * pangkat(n-1)
	}
}
