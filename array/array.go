package main

import "fmt"

func main() {
	type (
		names          [2]string // deklarasi tipe baru names sebagai array dengan panjang 2 dan tipe data string
		tempat_tinggal [3]string // deklarasi tipe baru tempat_tinggal dengan tipe dasar string
	)

	var nama names    // deklarasi variabel nama dengan tipe data names
	nama[0] = "adit"  // mengisi index ke 0 dengan nilai "adit"
	nama[1] = "putra" // mengisi index ke 1 dengan nilai "putra"

	fmt.Println(nama[0], nama[1]) // mencetak nilai index ke 0 dan 1

	fmt.Println(nama) // mencetak seluruh isi array a

	kota := tempat_tinggal{"jakarta", "bandung", "surabaya"} // deklarasi variabel tempat_tinggal dengan tipe data tempat_tinggal
	fmt.Println(kota[2])                                     // mencetak index ke 2 isi array kota
	fmt.Println(kota)                                        // mencetak seluruh isi array kota

	primes := [6]int{2, 3, 5, 7, 11, 13} // deklarasi array dengan inisialisasi nilai
	fmt.Println(primes)                  // mencetak seluruh isi array primes
	fmt.Println(primes[0] > primes[1])   // mencetak nilai index ke 0 dari array primes lebih besar dari index ke 1
}
