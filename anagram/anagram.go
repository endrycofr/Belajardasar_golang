package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Fungsi untuk sort huruf
func sortString(s string) string { // fungsi mengurutkan huruf dalam string
	sArr := strings.Split(s, "")  // memecah string menjadi slice karakter
	sort.Strings(sArr)            // mengurutkan slice karakter
	return strings.Join(sArr, "") // menggabungkan kembali menjadi string
}

func main() {
	reader := bufio.NewReader(os.Stdin) // untuk membaca input dari user

	fmt.Print("Masukkan kata-kata (pisahkan dengan koma): ") // prompt input
	inputString, _ := reader.ReadString('\n')                // baca input hingga newline
	inputString = strings.TrimSpace(inputString)             // bersihkan spasi di awal dan akhir

	// buat slice dari input
	words := strings.Split(inputString, ",") // pisahkan berdasarkan koma

	// bersihkan spasi tiap kata
	for i := range words { // iterasi tiap kata
		words[i] = strings.TrimSpace(words[i]) // bersihkan spasi di awal dan akhir
	}
	fmt.Println("Kata-kata yang dimasukkan:", words)
	// Map untuk kelompok anagram
	groups := make(map[string][]string) // key: string terurut, value: slice kata-kata

	// Kelompokkan kata-kata berdasarkan anagram

	for _, word := range words { // iterasi tiap kata
		key := sortString(word)                 // dapatkan key dengan mengurutkan huruf
		groups[key] = append(groups[key], word) // tambahkan kata ke kelompok yang sesuai
	}
	fmt.Println("Kelompok anagram:", groups)
	// Konversi ke slice 2 dimensi
	var result [][]string          // slice 2 dimensi untuk hasil akhir
	for _, group := range groups { // iterasi tiap kelompok
		result = append(result, group) // tambahkan kelompok ke hasil akhir
	}

	fmt.Println(result) // cetak hasil akhir
}
