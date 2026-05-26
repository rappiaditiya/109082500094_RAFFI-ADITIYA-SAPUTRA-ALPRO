package main

import "fmt"

func cetakDeret(n int) {
	for {
		fmt.Printf("%d ", n)
		if n == 1 {
			break
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}

func main() {
	var angka int
	fmt.Print("Masukkan angka awal: ")
	fmt.Scan(&angka)
	cetakDeret(angka)
}
