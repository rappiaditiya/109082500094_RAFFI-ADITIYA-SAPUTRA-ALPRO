package main

import "fmt"

// Fungsi untuk menghitung dan menampilkan rincian biaya pengiriman
func BiayaPos(beratGram int) {
	// Konversi berat ke kilogram dan sisa gramnya
	beratKg := beratGram / 1000  // Bagian berat dalam kg
	sisaGram := beratGram % 1000 // Sisa berat yang kurang dari 1 kg

	// Biaya pokok: setiap kg dikenakan biaya Rp10.000
	biayaKg := beratKg * 10000

	var tambahanBiaya int
	// Aturan perhitungan biaya tambahan untuk sisa gram:
	if beratKg >= 10 {
		tambahanBiaya = 0 // Jika berat total >= 10 kg, sisa gram GRATIS
	} else if sisaGram > 500 {
		tambahanBiaya = sisaGram * 5 // Jika sisa > 500g, per gram Rp5
	} else {
		tambahanBiaya = sisaGram * 15 // Jika sisa <= 500g, per gram Rp15
	}

	// Jumlahkan biaya pokok dan biaya tambahan
	totalBiaya := biayaKg + tambahanBiaya

	// Tampilkan rincian perhitungan
	fmt.Println("\n===== Detail Perhitungan =====")
	fmt.Printf("Detail berat  : %d kg + %d  gram\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya  : Rp. %d  + Rp. %d\n", biayaKg, tambahanBiaya)
	fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
}

func main() {
	var beratGram int
	// Minta pengguna masukkan berat dalam satuan gram
	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&beratGram)
	// Panggil fungsi hitung biaya
	BiayaPos(beratGram)
}
