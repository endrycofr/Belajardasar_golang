package main

import "fmt"

func main() {
	name := "Rizqi Fauzi"
	ages := []int{20, 21, 22, 23, 24}
	counter := 0

	// closure function untuk mengincrement counter
	increment := func() {
		name := "budi" // mengubah variabel dari scope luar
		fmt.Println("Nama di dalam closure increment:", name)
		counter++
		fmt.Println("Counter:", counter)
	}

	// closure function untuk menghitung panjang nama
	nameLength := func() {
		name := "Rizqi Fauzi Updated" // mengubah variabel dari scope luar
		fmt.Println("Nama di dalam closure:", name)
	}

	// closure function untuk menghitung jumlah elemen dalam slice ages
	ageCount := func() {
		ages := append(ages, 25) // mengubah slice dari scope luar
		fmt.Println("Nama di dalam closure:", ages)
	}
	// memanggil closure function
	nameLength()
	ageCount()
	increment()
	increment()

	// menampilkan nilai variabel dari scope luar setelah pemanggilan closure
	fmt.Println("Counter setelah pemanggilan increment:", counter)
	fmt.Println("Nama di luar closure:", name)
	fmt.Println("Ages di luar closure:", ages)
}
