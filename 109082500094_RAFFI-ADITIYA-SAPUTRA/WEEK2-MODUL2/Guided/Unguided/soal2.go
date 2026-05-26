package main

import "fmt"

func main() {
	// Mendefinisikan array berisi urutan warna standar yang harus diingat
	warna := [4]string{"merah", "kuning", "hijau", "ungu"}
	// Variabel penanda, diatur awalnya sebagai berhasil (true)
	berhasil := true

	// Melakukan perulangan sebanyak 5 kali percobaan
	for percobaan := 1; percobaan <= 5; percobaan++ {
		fmt.Printf("Percobaan %d : ", percobaan)

		// Menampung 4 input warna dari pengguna
		var warna1, warna2, warna3, warna4 string
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)

		// Mengecek apakah urutan input sama persis dengan urutan asli
		// Jika ada satu saja yang salah, ubah status menjadi gagal (false)
		if warna1 != warna[0] || warna2 != warna[1] ||
			warna3 != warna[2] || warna4 != warna[3] {
			berhasil = false
		}
	}
	// Menampilkan hasil akhir: berhasil hanya jika SEMUA percobaan benar
	fmt.Printf("BERHASIL : %t\n", berhasil)
}
