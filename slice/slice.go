package main

import "fmt"

func main() {
	slice()
	manipulasiSlice()
	makeslice()
}

func slice() {
	type nama []string
	namaSaya := nama{
		"Adit",
		"Putra",
		"Santoso",
		"Ganteng",
		"Sekali",
		"bryan",
		"park",
	}
	fmt.Println("daftar nama:", namaSaya[1:4]) // mencetak index ke 1 sampai 4 dari array namaSaya
	fmt.Println("panjang:", len(namaSaya))     // mencetak panjang dari array namaSaya
	fmt.Println("kapasitas:", cap(namaSaya))   // mencetak kapasitas dari array namaSaya

}
func manipulasiSlice() {
	type lahir [6]int
	tahun := lahir{1995, 1996, 3000, 1998, 1999, 2000}

	fmt.Println("tahun:", tahun[1:2])           // mencetak index ke 0 sampai 2 dari array tahun
	tahun[2] = 1997                             // mengubah index ke 0 menjadi 2000
	fmt.Println("tahun setelah diubah:", tahun) // mencetak seluruh isi array tahun setelah diubah index ke 0 menjadi 2000
	fmt.Println("panjang:", len(tahun))         // mencetak panjang dari array tahun
	fmt.Println("kapasitas:", cap(tahun))       // mencetak kapasitas dari array tahun

	tahunSlice := tahun[1:5]                         // membuat slice dari array tahun dengan index 1 sampai 4
	fmt.Println("tahun slice:", tahunSlice)          // mencetak seluruh isi slice tahunSlice
	fmt.Println("panjang slice:", len(tahunSlice))   // mencetak panjang dari slice tahunSlice
	fmt.Println("kapasitas slice:", cap(tahunSlice)) // mencetak kapasitas dari slice tahunSlice

	tahunBaru := append(tahunSlice, 2001)             // menambahkan nilai 1998 ke dalam slice hasil dari array tahun
	fmt.Println("tahun setelah ditambah:", tahunBaru) // mencetak seluruh isi slice setelah ditambahkan nilai 1998
	fmt.Println("tahun semua :", tahun)               // mencetak seluruh isi array tahun setelah slice ditambahkan nilai 1998

}

func makeslice() {
	tahunBaru2 := make([]int, 2, 5) // membuat slice baru dengan panjang 2 dan kapasitas 5
	tahunBaru2[0] = 2020            // mengisi index ke 0
	tahunBaru2[1] = 2021
	fmt.Println("tahun baru 2:", tahunBaru2)                      // mengisi index ke 1
	tahunBaru2 = append(tahunBaru2, 2022, 2023, 2024, 2025, 2026) // menambahkan nilai ke dalam slice tahunBaru2
	fmt.Println("tahun baru 2 ditambahkan:", tahunBaru2)          // mencetak seluruh isi slice tahunBaru2
	fmt.Println("panjang tahun baru 2:", len(tahunBaru2))
	fmt.Println("kapasitas tahun baru 2:", cap(tahunBaru2))

	copySlice := make([]int, len(tahunBaru2), cap(tahunBaru2)) // membuat slice baru dengan panjang sama dengan tahunBaru2
	copy(copySlice, tahunBaru2)                                // menyalin isi dari slice tahunBaru2 ke dalam slice copySlice
	fmt.Println("copy slice:", copySlice)                      // mencetak seluruh isi slice copySlice
}
