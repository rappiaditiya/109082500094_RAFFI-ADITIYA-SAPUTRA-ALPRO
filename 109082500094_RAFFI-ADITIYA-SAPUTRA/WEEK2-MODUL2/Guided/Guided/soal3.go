package main // Menandakan ini adalah program yang bisa dijalankan langsung
import "fmt" // Mengimpor pustaka untuk input dan output data

func main() {
	var usia, harga int // Deklarasi variabel: usia penonton dan harga tiket

	// Meminta pengguna memasukkan usia
	fmt.Print("Masukkan usia penonton: ")
	fmt.Scanln(&usia)

	// Menentukan harga tiket berdasarkan kategori usia
	if usia < 12 {
		harga = 30000 // Usia di bawah 12 tahun: harga Rp30.000
	} else if usia >= 12 && usia <= 17 {
		harga = 40000 // Usia 12 sampai 17 tahun: harga Rp40.000
	} else {
		harga = 50000 // Usia 18 tahun ke atas: harga Rp50.000
	}

	// Menampilkan hasil akhir
	fmt.Printf("Usia penonton: %d tahun\n", usia)
	fmt.Printf("Harga tiket yang harus dibayar: Rp%d\n", harga)
}
