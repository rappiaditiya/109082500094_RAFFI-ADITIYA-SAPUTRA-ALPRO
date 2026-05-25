package main

import (
	"fmt"
	"math"
)

// Prosedur hitung persegi
func hitungPersegi(sisi int) {
	luas := sisi * sisi
	keliling := 4 * sisi
	fmt.Printf("Luas persegi : %d\n", luas)
	fmt.Printf("Keliling persegi : %d\n", keliling)
}

// Prosedur hitung persegi panjang
func hitungPersegiPanjang(panjang, lebar int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	fmt.Printf("Luas persegi panjang : %d\n", luas)
	fmt.Printf("Keliling persegi panjang : %d\n", keliling)
}

// Prosedur hitung lingkaran
func hitungLingkaran(jarijari float64) {
	luas := math.Pi * jarijari * jarijari
	keliling := 2 * math.Pi * jarijari
	fmt.Printf("Luas lingkaran : %.5f\n", luas)
	fmt.Printf("Keliling lingkaran : %.5f\n", keliling)
}

func main() {
	var pilihan int
	var sisi, panjang, lebar int
	var jarijari float64

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print("Masukan sisi : ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)
	case 2:
		fmt.Print("Masukan panjang : ")
		fmt.Scan(&panjang)
		fmt.Print("Masukan lebar : ")
		fmt.Scan(&lebar)
		hitungPersegiPanjang(panjang, lebar)
	case 3:
		fmt.Print("Masukan jari-jari : ")
		fmt.Scan(&jarijari)
		hitungLingkaran(jarijari)
	default:
		fmt.Println("Pilihan tidak tersedia!")
	}
}
