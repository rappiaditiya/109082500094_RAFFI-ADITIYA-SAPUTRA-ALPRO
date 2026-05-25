package main

import "fmt"

func ganjil(sampai, sekarang int) {
	if sekarang <= sampai {
		if sekarang%2 != 0 {
			fmt.Printf("%d ", sekarang)
		}
		ganjil(sampai, sekarang+1)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	ganjil(N, 1)
}
