package main

import "fmt"

func main() {
	suara := make([]int, 21)
	totalTerbaca := 0
	totalSah := 0
	var angka int

	for {
		fmt.Scan(&angka)
		if angka == 0 {
			break
		}
		totalTerbaca++

		if angka >= 1 && angka <= 20 {
			suara[angka]++
			totalSah++
		}
	}

	fmt.Println("Suara masuk:", totalTerbaca)
	fmt.Println("Suara sah:", totalSah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}
