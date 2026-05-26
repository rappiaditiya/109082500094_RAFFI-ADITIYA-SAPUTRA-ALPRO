package main

import "fmt"

// Definisi struktur data mahasiswa
type mahasiswa struct {
	nama string
	nim  string
	IPK  float64
}

// Tipe data array mahasiswa dengan ukuran tetap 3
type arrMahasiswa [3]mahasiswa

// Fungsi mencari mahasiswa dengan IPK terbesar
func IPKTerbesar(array []mahasiswa) (mahasiswa, int) {
	var mhsTerbesar mahasiswa = array[0]
	var idxDitemukan int = 0
	var i int // Diperbaiki: huruf kecil agar konsisten

	for i = 1; i < len(array); i++ {
		if array[i].IPK > mhsTerbesar.IPK {
			mhsTerbesar = array[i]
			idxDitemukan = i
		}
	}
	return mhsTerbesar, idxDitemukan
}

// Fungsi mencari mahasiswa dengan IPK terkecil
func IPKTerkecil(array []mahasiswa) (mahasiswa, int) {
	var mhsTerkecil mahasiswa = array[0]
	var idxDitemukan int = 0
	var i int // Diperbaiki: huruf kecil agar konsisten

	for i = 1; i < len(array); i++ {
		if array[i].IPK < mhsTerkecil.IPK {
			mhsTerkecil = array[i]
			idxDitemukan = i
		}
	}
	return mhsTerkecil, idxDitemukan
}

func main() {
	var arrMhs arrMahasiswa
	var i, j int // Diperbaiki: huruf kecil agar konsisten

	// Menginput data mahasiswa
	for i = 0; i < len(arrMhs); i++ {
		fmt.Println("DATA MAHASISWA INDEKS KE-", i)
		fmt.Print("Masukkan nama : ")
		fmt.Scan(&arrMhs[i].nama)
		fmt.Print("Masukkan NIM : ")
		fmt.Scan(&arrMhs[i].nim)
		fmt.Print("Masukkan IPK : ")
		fmt.Scan(&arrMhs[i].IPK)
		fmt.Println("========================")
	}

	// Menampilkan kembali seluruh data yang diinput
	fmt.Println()
	for j = 0; j < len(arrMhs); j++ {
		fmt.Println("DATA MAHASISWA INDEKS KE-", j)
		fmt.Println("Nama : ", arrMhs[j].nama)
		fmt.Println("NIM : ", arrMhs[j].nim)
		fmt.Println("IPK : ", arrMhs[j].IPK)
		fmt.Println("========================")
	}

	// Memanggil fungsi pencarian
	var mhsMaks, mhsMin mahasiswa
	var idxMaks, idxMin int

	mhsMaks, idxMaks = IPKTerbesar(arrMhs[:])
	mhsMin, idxMin = IPKTerkecil(arrMhs[:])

	// Menampilkan hasil pencarian
	fmt.Println()
	fmt.Println("=== DATA IPK TERKECIL & TERBESAR MAHASISWA ===")
	fmt.Printf("IPK Terbesar : %.2f, atas nama %s dengan NIM %s, ditemukan pada indeks ke-%d\n",
		mhsMaks.IPK, mhsMaks.nama, mhsMaks.nim, idxMaks)
	fmt.Printf("IPK Terkecil : %.2f, atas nama %s dengan NIM %s, ditemukan pada indeks ke-%d\n",
		mhsMin.IPK, mhsMin.nama, mhsMin.nim, idxMin)
}
