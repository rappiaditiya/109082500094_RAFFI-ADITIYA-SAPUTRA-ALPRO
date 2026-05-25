package main

import "fmt"

const NMAX = 1000000  // jumlah maks data masukan
type arrInt [NMAX]int // tipe data alias array integer

func SelectionSort(T arrInt, n int) {
	/* I.S. terdefinisi array T yang berisi sejumlah n bilangan bulat
	   F.S. array T terurut secara membesar berdasarkan algoritma selection sort */
	var idxMin, temp int
	for i := 0; i < n-1; i++ {
		idxMin = i
		for j := i + 1; j < n; j++ {
			if T[j] < T[idxMin] {
				idxMin = j
			}
		}
		// tukar elemen
		temp = T[i]
		T[i] = T[idxMin]
		T[idxMin] = temp
	}
}

func median(T arrInt, n int) float64 {
	/* mengembalikan median dari array T yang berisi sejumlah n bilangan bulat
	   yang telah terurut membesar */
	if n%2 == 1 {
		// jumlah ganjil: ambil nilai tengah
		return float64(T[n/2])
	} else {
		// jumlah genap: rata-rata dua nilai tengah, dibulatkan ke bawah
		tengah1 := T[(n/2)-1]
		tengah2 := T[n/2]
		return float64((tengah1 + tengah2) / 2)
	}
}

func main() {
	var A arrInt  // array integer
	var x int     // variabel masukan
	var n int = 0 // jumlah data yang tersimpan

	fmt.Print("Input data masukan : ") // sesuai teks di contoh keluaran
	fmt.Scan(&x)
	for x != -5313541 && n < NMAX {
		if x == 0 {
			// jika ketemu angka 0: urutkan lalu hitung & cetak median
			SelectionSort(A, n)
			fmt.Println("Median :")
			fmt.Printf("%.1f\n", median(A, n))
		} else {
			// simpan data selain 0 dan penanda akhir
			A[n] = x
			n++
		}
		fmt.Scan(&x)
	}
}
