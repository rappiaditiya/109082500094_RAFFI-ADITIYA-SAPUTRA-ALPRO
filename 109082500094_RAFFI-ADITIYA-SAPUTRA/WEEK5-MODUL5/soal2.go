package main

import "fmt"

func cetakBintang(jumlah int) {
	if jumlah > 0 {
		cetakBintang(jumlah - 1)
		fmt.Print("*")
	}
}

func cetakBaris(n, saatIni int) {
	if saatIni <= n {
		cetakBintang(saatIni)
		fmt.Println()
		cetakBaris(n, saatIni+1)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	cetakBaris(N, 1)
}
