package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var data []int
	var angka int
	for {
		fmt.Scan(&angka)
		if angka == -5313 {
			break
		}
		if angka == 0 {

			insertionSort(data)
			n := len(data)
			if n%2 == 1 {
				// Ganjil
				fmt.Println(data[n/2])
			} else {

				tengah1 := data[(n/2)-1]
				tengah2 := data[n/2]
				fmt.Println((tengah1 + tengah2) / 2)
			}
		} else {

			data = append(data, angka)
		}
	}
}
