package main

import "fmt"

const MAX = 1000

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	var berat [MAX]float64

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := (x + y - 1) / y
	totalBerat := make([]float64, jumlahWadah)
	rataRata := make([]float64, jumlahWadah)

	for i := 0; i < x; i++ {
		indeksWadah := i / y
		totalBerat[indeksWadah] += berat[i]
	}

	for i := 0; i < jumlahWadah; i++ {
		if i < jumlahWadah-1 || x%y == 0 {
			rataRata[i] = totalBerat[i] / float64(y)
		} else {
			sisa := x % y
			if sisa == 0 {
				sisa = y
			}
			rataRata[i] = totalBerat[i] / float64(sisa)
		}
	}

	for i := 0; i < jumlahWadah; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", totalBerat[i])
	}
	fmt.Println()

	for i := 0; i < jumlahWadah; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", rataRata[i])
	}
	fmt.Println()
}
