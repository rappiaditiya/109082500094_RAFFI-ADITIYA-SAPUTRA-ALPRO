package main

import "fmt"

func cetakUrutan(n, saatIni int) {
	if saatIni >= 1 {
		fmt.Printf("%d ", saatIni)
		cetakUrutan(n, saatIni-1)
	}
	if saatIni < n && saatIni > 0 {
		fmt.Printf("%d ", saatIni+1)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	cetakUrutan(N, N)
}
