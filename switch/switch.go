package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() { // contoh penggunaan switch
	reader := bufio.NewReader(os.Stdin)  // membaca input dari user
	fmt.Print("Masukkan sebuah angka: ") // prompt input
	input, _ := reader.ReadString('\n')  // baca input hingga newline
	input = strings.TrimSpace(input)     // bersihkan spasi di awal dan akhir
	// konversi string ke int
	angka, err := strconv.Atoi(input) // konversi string ke int
	if err != nil {                   // cek error konversi
		fmt.Println("Error membaca input:", err) // tampilkan pesan error
		return                                   // keluar dari program
	}
	switch sum := angka; { // switch tanpa ekspresi
	case sum >= 80: // jika sum >= 80
		println("Nilai A") // tampilkan A
	case sum >= 70: // jika sum >= 70
		println("Nilai B") // tampilkan B
	case sum >= 60: // jika sum >= 60
		println("Nilai C") // tampilkan C
	default: // jika tidak ada yang cocok
		println("Nilai D") // tampilkan D
	}
}
