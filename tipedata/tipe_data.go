package main

import "fmt"

//tipe data dasar: integer, float, boolean, string
func main() {
	fmt.Println("lima =", 5)             //integer
	fmt.Println("sepuluh =", 10)         //integer
	fmt.Println("tiga koma lima =", 3.5) //float
	fmt.Println("true =", true)          //boolean
	fmt.Println("false =", false)        //boolean
	fmt.Println("hello =", "hello")      //string

	fmt.Println(len("rizqi")) //menghitung panjang string "rizqi"
	fmt.Println("rizqi"[0])   //mengambil karakter ke 0 dari string "rizqi"
	fmt.Println("rizqi"[2])   //mengambil karakter ke 2 dari string "rizqi"

	fmt.Println(len("rizqi" + " " + "fauzi")) //menggabungkan string "rizqi" dan "fauzi" lalu menghitung panjangnya
}
