package main

import "fmt"

// Fungsi rekursif untuk menghitung nilai Fibonacci ke-n
func fibonacci(n int) int {
	// Kondisi dasar: suku ke-0 bernilai 0
	if n == 0 {
		return 0
		// Kondisi dasar: suku ke-1 bernilai 1
	} else if n == 1 {
		return 1
	}
	// Suku selanjutnya = jumlah dari dua suku sebelumnya
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Deret Fibonacci hingga suku ke-10: ")
	// Mencetak judul kolom dengan format rata kiri dan lebar 5 karakter
	fmt.Printf("%-5s %-5s\n", "n", "Sn")

	// Perulangan dari suku ke-0 sampai ke-10
	for i := 0; i <= 10; i++ {
		// Menampilkan nomor suku dan nilainya
		fmt.Printf("%-5d %-5d\n", i, fibonacci(i))
	}
}
