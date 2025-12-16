package main

import (
	"fmt"
	"strings"
)

type name struct { // mendefinisikan struct name
	FirstName string
	LastName  string
}

func GetFullName(n name) string { // fungsi untuk menggabungkan nama depan dan belakang
	return n.FirstName + " " + n.LastName
}
func main() {
	n := name{FirstName: "John Dane", LastName: "John Doe "} // membuat instance struct name
	fullName := GetFullName(n)                               // memanggil fungsi untuk mendapatkan nama lengkap

	// Demonstrasi beberapa fungsi dari package strings
	fmt.Println("Full Name:", fullName)
	fmt.Println("Contains search 'johan':", strings.Contains(fullName, "Johan")) // contains bersifat case-sensitive karena huruf besar dan kecil dianggap berbeda
	fmt.Println("Split by space:", strings.Split(fullName, " "))                 // memisahkan string berdasarkan spasi
	fmt.Println("To Lower Case:", strings.ToLower(fullName))                     // mengubah semua karakter menjadi huruf kecil
	fmt.Println("To Upper Case:", strings.ToUpper(fullName))                     // mengubah semua karakter menjadi huruf besar
	fmt.Println("Trimmed spaces:", strings.Trim(fullName, " "))                  // menghapus spasi di awal dan akhir string

	fmt.Println("Replace 'John' with 'Jane':", strings.Replace(fullName, "John", "Jane", -1)) // mengganti 'John' dengan 'Jane'
	// argumen terakhir '1' menunjukkan hanya mengganti kemunculan pertama
	// argumen ini bisa diisi -1 untuk mengganti semua kemunculan
	// argumen ini bisa diisi 0 untuk tidak mengganti sama sekali
	// argumen ini bisa diisi n untuk mengganti sebanyak n kemunculan
	//(bersifat case-sensitive karena huruf besar dan kecil dianggap berbeda)
}
