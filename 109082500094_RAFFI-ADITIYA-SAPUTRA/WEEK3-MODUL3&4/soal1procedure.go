package main

import "fmt"

// Prosedur buat ngecetak deretnya
func cetakDeret(n int) {
	for {
		fmt.Printf("%d ", n)
		// Berhenti kalau udah sampai angka 1
		if n == 1 {
			break
		}
		// Aturan hitung: genap dibagi 2, ganjil jadi 3n+1
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}

func main() {
	var awal int
	// Tulisan perintah masukin angka
	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&awal)
	cetakDeret(awal)
}
