package main

import "fmt"

type arrBalita [100]float64

// Fungsi mencari nilai minimum dan maksimum, hasil dikembalikan lewat parameter pointer
func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Fungsi menghitung rata-rata berat
func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var min, max float64

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	// Membatasi jumlah data maksimal 100
	if n > 100 {
		n = 100
	}

	// Menginput data berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	// Memanggil fungsi untuk cari min dan max (kirim alamat variabel min & max)
	hitungMinMax(data, n, &min, &max)
	// Memanggil fungsi hitung rata-rata
	rata := rerata(data, n)

	// Menampilkan hasil akhir
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata) // <-- Diperbaiki: 'kh' menjadi 'kg'
}
