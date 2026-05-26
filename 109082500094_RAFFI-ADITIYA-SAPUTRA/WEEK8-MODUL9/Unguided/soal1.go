package main

import (
	"fmt"
	"math"
)

// Titik mendefinisikan koordinat x dan y
type Titik struct {
	x, y float64
}

// Lingkaran memiliki titik pusat dan jari-jari
type Lingkaran struct {
	pusat Titik
	r     float64
}

// jarak menghitung jarak antara dua titik menggunakan rumus jarak Euclidean
func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

// didalam mengecek apakah titik p berada di dalam atau di batas lingkaran c
func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= c.r
}

func main() {
	var l1, l2 Lingkaran
	var p Titik

	// Membaca input: pusat x, pusat y, jari-jari untuk lingkaran 1
	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	// Membaca input: pusat x, pusat y, jari-jari untuk lingkaran 2
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	// Membaca input: koordinat titik yang akan diperiksa
	fmt.Scan(&p.x, &p.y)

	// Cek posisi titik terhadap kedua lingkaran
	diL1 := didalam(l1, p)
	diL2 := didalam(l2, p)

	// Menentukan dan menampilkan hasil
	if diL1 && diL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
