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

	suaraTertinggi1 := -1
	ketua := -1
	for i := 1; i <= 20; i++ {
		if suara[i] > suaraTertinggi1 {
			suaraTertinggi1 = suara[i]
			ketua = i
		}
	}

	suaraTertinggi2 := -1
	wakil := -1
	for i := 1; i <= 20; i++ {
		if i == ketua {
			continue
		}
		if suara[i] > suaraTertinggi2 {
			suaraTertinggi2 = suara[i]
			wakil = i
		}
	}

	fmt.Println("Suara masuk:", totalTerbaca)
	fmt.Println("Suara sah:", totalSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
