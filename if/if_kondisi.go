package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)  // membaca input dari user
	fmt.Print("Masukkan sebuah angka: ") // prompt input
	input, _ := reader.ReadString('\n')  // baca input hingga newline

	input = strings.TrimSpace(input) // bersihkan spasi di awal dan akhir
	// konversi string ke int
	angka, err := strconv.Atoi(input) // konversi string ke int
	if err != nil {                   // cek error konversi
		fmt.Println("Error membaca input:", err) // tampilkan pesan error
		return                                   // keluar dari program
	} // cek kondisi genap atau ganjil
	if angka%2 == 0 { // jika angka habis dibagi 2
		println("Angka adalah genap") // tampilkan genap

	} else {
		println("Angka adalah ganjil") // tampilkan ganjil
	}
}
