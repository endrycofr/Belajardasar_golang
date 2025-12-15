package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi cek prima (clean)
func prima(n int) bool { // fungsi untuk mengecek bilangan prima
	if n <= 1 { // bilangan prima adalah bilangan > 1
		return false // bukan prima
	}
	for i := 2; i*i < n; i++ { // cek pembagi dari 2 hingga akar n
		if n%i == 0 { // jika n habis dibagi i
			return false // bukan prima
		}
	}
	return true // prima
}

// Fungsi cek ganjil / genap
func bilangan(n int) string { // fungsi untuk mengecek ganjil atau genap
	if n%2 == 0 { // jika n habis dibagi 2
		return "genap" // genap
	}
	return "ganjil" // ganjil
}

func main() {
	reader := bufio.NewReader(os.Stdin) // membaca input dari user

	fmt.Print("Masukkan sebuah angka: ") // prompt input
	input, _ := reader.ReadString('\n')  // baca input hingga newline
	input = strings.TrimSpace(input)     // hapus spasi di awal dan akhir

	angka, err := strconv.Atoi(input) // konversi string ke int
	if err != nil {                   // cek error konversi
		fmt.Println("Input bukan angka!") // tampilkan pesan error
		return
	}

	// Panggil sekali saja
	jenis := bilangan(angka) // cek ganjil/genap
	isPrima := prima(angka)  // cek prima

	// Output final
	if isPrima { // jika prima
		fmt.Printf("Angka %d adalah bilangan %s dan bilangan prima\n", angka, jenis) // tampilkan hasil
	} else { // jika bukan prima
		fmt.Printf("Angka %d adalah bilangan %s dan bukan bilangan prima\n", angka, jenis) // tampilkan hasil
	}

	// Panggil berulang kali
	namasake() // panggil fungsi namasake
	println("menggunakan range untuk mengambil slice")
	namerange() // panggil fungsi namerange
}

func namasake() string {
	slice := []string{"Andi", "Budi", "Caca", "Dodi", "Eka"} // slice nama-nama
	for i := 0; i < len(slice); i++ {                        // iterasi menggunakan for biasa
		println("nama ke", i+1, "adalah", slice[i]) // tampilkan nama

	}
	return "" // kembalikan string kosong

}
func namerange() string { // fungsi menggunakan range
	slice := []string{"Andi", "Budi", "Caca", "Dodi", "Eka"}
	for i, v := range slice { // iterasi menggunakan range untuk mendapatkan index dan value
		println("nama ke", i+1, "adalah", v)
	}

	return "" // kembalikan string kosong
}
