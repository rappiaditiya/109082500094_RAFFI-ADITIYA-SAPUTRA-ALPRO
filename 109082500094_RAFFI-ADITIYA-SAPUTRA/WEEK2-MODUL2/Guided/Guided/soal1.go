package main // Menandakan ini adalah program yang bisa dieksekusi, bukan pustaka

import "fmt" // Mengimpor pustaka fmt untuk fungsi input/output

func main() { // Titik awal eksekusi program
	// Mendeklarasikan variabel berat, tinggi, dan bmi dengan tipe data float64 (bilangan desimal)
	var berat, tinggi, bmi float64

	// Meminta pengguna memasukkan berat badan dan menyimpannya ke variabel berat
	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scanln(&berat)

	// Meminta pengguna memasukkan tinggi badan dan menyimpannya ke variabel tinggi
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scanln(&tinggi)

	// Rumus perhitungan BMI: berat dibagi kuadrat tinggi badan
	bmi = berat / (tinggi * tinggi)

	// Menampilkan hasil BMI dengan 2 angka di belakang koma
	fmt.Printf("Nilai BMI Anda adalah: %.2f\n", bmi)
}
