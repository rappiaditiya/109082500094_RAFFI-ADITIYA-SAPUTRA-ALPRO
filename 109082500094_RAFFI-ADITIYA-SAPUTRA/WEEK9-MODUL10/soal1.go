package main

import "fmt"

const MAX = 1000

func main() {
	var n int
	fmt.Scan(&n)
	var berat [MAX]float64

	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]
	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Printf("%.2f %.2f\n", min, max)
}
