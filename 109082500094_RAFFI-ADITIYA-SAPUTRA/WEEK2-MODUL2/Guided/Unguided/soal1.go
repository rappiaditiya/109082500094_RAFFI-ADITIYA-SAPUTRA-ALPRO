package main // Menandakan ini adalah program yang bisa dijalankan langsung
import "fmt" // Mengimpor pustaka untuk input dan output data

// Fungsi untuk mengecek tahun kabisat, mengembalikan nilai true atau false
func cekKabisat(tahun int) bool {
	// Aturan 1: Tahun yang habis dibagi 400 pasti kabisat
	if tahun%400 == 0 {
		return true
		// Aturan 2: Tahun yang habis dibagi 4, TAPI tidak habis dibagi 100, itu kabisat
	} else if tahun%4 == 0 && tahun%100 != 0 {
		return true
	}
	// Selain dari aturan di atas, BUKAN tahun kabisat
	return false
}

func main() {
	var tahun int
	// Meminta pengguna memasukkan angka tahun
	fmt.Print("Masukkan tahun : ")
	fmt.Scan(&tahun)
	// Memanggil fungsi cekKabisat dan menampilkan hasilnya (true/false)
	fmt.Printf("Kabisat : %t\n", cekKabisat(tahun))
}
