package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, idxHapus, cariFrekuensi int
	const kapasitas = 100
	var arr [kapasitas]int

	fmt.Print("Masukkan jumlah elemen (N): ")
	fmt.Scan(&n)

	fmt.Printf("Masukkan %d elemen angka:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// a. Isi seluruh array
	fmt.Print("\na. Isi seluruh array: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	// b. Elemen pada indeks ganjil
	fmt.Print("b. Elemen indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	// c. Elemen pada indeks genap
	fmt.Print("c. Elemen indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	// d. Elemen pada indeks kelipatan x
	fmt.Print("d. Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()

	// e. (Kosong di kode asli, lanjut ke f)

	// f. Rata-rata
	var total float64
	for i := 0; i < n; i++ {
		total += float64(arr[i])
	}
	rataRata := total / float64(n)
	fmt.Printf("f. Rata-rata: %.2f\n", rataRata)

	// g. Standar Deviasi
	var jumlahKuadratSelisih float64
	for i := 0; i < n; i++ {
		jumlahKuadratSelisih += math.Pow(float64(arr[i])-rataRata, 2)
	}
	stdDev := math.Sqrt(jumlahKuadratSelisih / float64(n))
	fmt.Printf("g. Standar Deviasi: %.2f\n", stdDev)

	// h. Frekuensi bilangan tertentu
	fmt.Print("h. Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cariFrekuensi)
	count := 0
	for i := 0; i < n; i++ {
		if arr[i] == cariFrekuensi {
			count++
		}
	}
	fmt.Printf("Frekuensi %d dalam array: %d kali\n", cariFrekuensi, count)

	// Tambahan: Contoh penggunaan variabel idxHapus yang sudah dideklarasikan
	fmt.Print("\ni. Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idxHapus)
	if idxHapus >= 0 && idxHapus < n {
		// Geser elemen ke kiri untuk menimpa elemen yang dihapus
		for i := idxHapus; i < n-1; i++ {
			arr[i] = arr[i+1]
		}
		n-- // Kurangi jumlah elemen aktif
		fmt.Print("Array setelah dihapus: ")
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", arr[i])
		}
		fmt.Println()
	} else {
		fmt.Println("Indeks tidak valid")
	}
}
