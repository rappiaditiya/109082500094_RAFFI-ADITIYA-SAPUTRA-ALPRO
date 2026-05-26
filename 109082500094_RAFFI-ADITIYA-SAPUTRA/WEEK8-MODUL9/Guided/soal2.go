package main

import "fmt"

func main() {
	// Deklarasi map bernama nilai dengan key bertipe string dan value bertipe integer
	var nilai map[string]int = make(map[string]int)

	// Menambahkan data ke dalam map
	nilai["Dhimas"] = 90
	nilai["Ichya"] = 90

	// Menampilkan seluruh isi map menggunakan perulangan for range
	fmt.Println("Data nilai : ")
	var nama string // variabel untuk menampung key
	var grade int   // variabel untuk menampung value

	for nama, grade = range nilai {
		fmt.Println(nama, " = ", grade)
	}

	// Operasi update: mengubah nilai untuk key "Ichya"
	nilai["Ichya"] = 80

	// Operasi searching: mencari data berdasarkan key
	var cariNama string = "hafizh" // Nama yang ingin dicari
	var isiData int                // Menampung nilai jika ditemukan
	var ok bool                    // Status pengecekan (true = ada, false = tidak ada)

	isiData, ok = nilai[cariNama]
	if ok {
		fmt.Println("Nilai", cariNama, " = ", isiData)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	// Operasi penghapusan: menghapus data dengan key "Dhimas"
	delete(nilai, "Dhimas")

	// Menampilkan data setelah penghapusan untuk memastikan hasilnya
	fmt.Println("\nData setelah dihapus:")
	for nama, grade = range nilai {
		fmt.Println(nama, " = ", grade)
	}
}
