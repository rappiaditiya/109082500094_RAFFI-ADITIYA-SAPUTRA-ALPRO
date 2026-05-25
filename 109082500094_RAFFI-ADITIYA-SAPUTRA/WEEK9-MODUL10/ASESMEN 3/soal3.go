package main

import "fmt"

const NMAX = 1000000

// struct partai
type partai struct {
	nama  int
	suara int
}

// tipe tabPartai: array of partai dengan kapasitas NMAX
type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	/* mengembalikan indeks partai yang memiliki nama yang dicari
	   pada array t yang berisi n partai atau -1 apabila tidak
	   ditemukan, gunakan sekuensial search */
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {
	// deklarasi variabel
	var p tabPartai
	var x, n, idx int

	fmt.Println("Masukan proses input suara:")

	// lakukan proses input suara secara berulang di sini, simpan
	// ke dalam array p, sehingga terdapat array p yang berisi hasil
	// peroleh suara n partai.
	n = 0
	fmt.Scan(&x)
	for x != -1 && n < NMAX {
		idx = posisi(p, n, x)
		if idx != -1 {
			// jika sudah ada, tambah suaranya
			p[idx].suara++
		} else {
			// jika belum ada, tambah partai baru
			p[n].nama = x
			p[n].suara = 1
			n++
		}
		fmt.Scan(&x)
	}

	// lakukan proses pengurutan dengan insertion sort descending
	// berdasarkan jumlah suara yang diperoleh.
	var temp partai
	for i := 1; i < n; i++ {
		temp = p[i]
		j := i - 1
		for j >= 0 && p[j].suara < temp.suara {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = temp
	}

	// tampilkan array p sesuai format
	fmt.Println("Hasil Perhitungan suara:")
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d(%d)", p[i].nama, p[i].suara)
	}
	fmt.Println()
}
