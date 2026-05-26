package main

import "fmt"

// Fungsi selection sort untuk mengurutkan nomor rumah membesar
func selectionSort(rumah []int) {
	var idxTerkecil, i, j int
	// Perulangan untuk menentukan posisi yang akan diisi
	for i = 0; i < len(rumah)-1; i++ {
		idxTerkecil = i // Awalnya anggap data di indeks i adalah yang terkecil
		// Cari data terkecil dari posisi i+1 sampai akhir
		for j = i + 1; j < len(rumah); j++ {
			if rumah[j] < rumah[idxTerkecil] {
				idxTerkecil = j // Simpan indeks data yang lebih kecil
			}
		}
		// Tukar posisi data terkecil ke posisi yang seharusnya
		if idxTerkecil != i {
			rumah[i], rumah[idxTerkecil] = rumah[idxTerkecil], rumah[i]
		}
	}
}

func main() {
	var n int
	// Membaca jumlah daerah
	fmt.Scan(&n)

	// Memproses setiap daerah
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m) // Membaca jumlah rumah di daerah tersebut

		// Membuat wadah untuk menampung nomor rumah
		daftarRumah := make([]int, m)

		// Membaca semua nomor rumah
		for j := 0; j < m; j++ {
			fmt.Scan(&daftarRumah[j])
		}

		// Mengurutkan nomor rumah menggunakan selection sort
		selectionSort(daftarRumah)

		// Menampilkan hasil urutan
		for k := 0; k < m; k++ {
			if k > 0 {
				fmt.Print(" ") // Beri spasi antar nomor
			}
			fmt.Print(daftarRumah[k])
		}
		fmt.Println() // Pindah baris untuk daerah berikutnya
	}
}
