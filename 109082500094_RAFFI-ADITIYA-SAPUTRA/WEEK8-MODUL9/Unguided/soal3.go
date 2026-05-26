package main

import "fmt"

func main() {
	// Deklarasi variabel (semua nama sudah konsisten)
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string // Menyimpan hasil setiap pertandingan

	// Input nama klub
	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)

	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	// Perulangan input skor
	i := 1
	for {
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&skorA, &skorB) // Memperbaiki penulisan variabel agar sama semua

		// Berhenti jika ada skor negatif
		if skorA < 0 || skorB < 0 {
			break
		}

		// Logika menentukan pemenang
		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}

		i++
	}

	// Menampilkan semua hasil
	fmt.Println()
	for indeks, hasil := range pemenang {
		fmt.Printf("Hasil %d : %s\n", indeks+1, hasil)
	}

	fmt.Println("Pertandingan selesai")
}
