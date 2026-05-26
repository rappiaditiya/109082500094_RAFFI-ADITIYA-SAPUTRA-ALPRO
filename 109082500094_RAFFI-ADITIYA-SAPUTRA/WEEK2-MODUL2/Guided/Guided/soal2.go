package main // Menandakan ini adalah program yang bisa dijalankan langsung
import "fmt" // Mengimpor pustaka untuk input dan output data

func main() {
	// Deklarasi variabel:
	// target = jumlah uang yang ingin dicapai
	// tabungan = uang yang ditabung setiap hari
	// total = akumulasi seluruh uang tabungan
	// hari = penghitung berapa hari sudah berlalu
	var target, tabungan, total, hari int

	// Meminta pengguna memasukkan jumlah target uang
	fmt.Print("Masukkan target uang yang ingin dicapai: ")
	fmt.Scanln(&target)

	// Inisialisasi nilai awal: total uang 0, hari dimulai dari 0
	total = 0
	hari = 0

	// Perulangan berjalan SELAMA total uang yang terkumpul BELUM mencapai target
	for total < target {
		hari++ // Tambah hitungan hari sebelum memasukkan nominal hari itu
		// Minta input nominal tabungan hari ini
		fmt.Printf("Masukkan nominal tabungan hari ke-%d: ", hari)
		fmt.Scanln(&tabungan)
		// Tambahkan nominal hari ini ke total keseluruhan
		total = total + tabungan
	}
	// Setelah perulangan berhenti (total >= target), tampilkan hasil
	fmt.Printf("Selamat! Target tercapai dalam %d hari.\n", hari)
	fmt.Printf("Total tabungan Anda terkumpul: Rp%d\n", total)
}
