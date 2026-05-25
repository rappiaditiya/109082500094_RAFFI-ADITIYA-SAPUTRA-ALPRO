package main

import "fmt"

const NMAX = 100

// Definisi tipe data pemain
type Pemain struct {
	nama   string
	gol    int
	assist int
}

// Definisi tipe array pemain
type TabPemain [NMAX]Pemain

func InsertionSortData(T *TabPemain, n int) {
	/* Mengurutkan data menurun:
	   1. Berdasarkan jumlah gol terbanyak
	   2. Jika gol sama, urut berdasarkan assist terbanyak */
	var temp Pemain
	var j int
	for i := 1; i < n; i++ {
		temp = T[i]
		j = i - 1
		// Syarat perbandingan:
		// - Jika gol sebelumnya lebih sedikit, geser
		// - Jika gol sama tapi assist sebelumnya lebih sedikit, geser
		for j >= 0 && (T[j].gol < temp.gol || (T[j].gol == temp.gol && T[j].assist < temp.assist)) {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = temp
	}
}

func main() {
	var data TabPemain
	var n int

	// Membaca masukan
	fmt.Println("Masukan data Input :")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i].nama, &data[i].gol, &data[i].assist)
	}

	// Proses pengurutan
	InsertionSortData(&data, n)

	// Menampilkan hasil
	fmt.Println("Hasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %d\n", data[i].nama, data[i].gol, data[i].assist)
	}
}
