package main

import "fmt"

func selectionSort(arr []int, m int) {
	for i := 0; i < m-1; i++ {
		minIdx := i
		for j := i + 1; j < m; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var m, x int
		fmt.Scan(&m)
		var ganjil, genap []int
		for j := 0; j < m; j++ {
			fmt.Scan(&x)
			if x%2 == 1 {
				ganjil = append(ganjil, x)
			} else {
				genap = append(genap, x)
			}
		}
		selectionSort(ganjil, len(ganjil))
		selectionSort(genap, len(genap))

		for idx := 0; idx < len(ganjil); idx++ {
			if idx > 0 {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[idx])
		}

		for idx := 0; idx < len(genap); idx++ {
			fmt.Print(" ")
			fmt.Print(genap[idx])
		}
		fmt.Println()
	}
}
