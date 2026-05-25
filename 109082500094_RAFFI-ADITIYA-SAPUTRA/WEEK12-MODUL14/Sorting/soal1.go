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
		var m int
		fmt.Scan(&m)
		data := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&data[j])
		}
		selectionSort(data, m)
		for j := 0; j < m; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(data[j])
		}
		fmt.Println()
	}
}
