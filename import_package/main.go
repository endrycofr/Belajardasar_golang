package main

//import package helper

import (
	"fmt"
	"import_package/database"
	"import_package/helper"
	// import blank identifier menambahkan _ sebelum nama package untuk mengimpor package hanya \
	// untuk menjalankan init function nya saja tanpa mengakses apapun dari package tersebut
	//_ "import_package/databasesql"
)

func main() {

	result := database.GetConnection() //memanggil fungsi GetConnection dari package database
	fmt.Println(result)                //menampilkan status koneksi database

	//===== import package helper =====
	c := helper.GetClass("Alice", 20, 1012)
	fmt.Println("Name:", c.Name)
	fmt.Println("Age:", c.Age)
	fmt.Println("NoID:", c.NoID)
	// fmt.Println("ID:", c.GetID()) //mengakses Id melalui method GetID karena field id tidak bisa diakses langsung

}
