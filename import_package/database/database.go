package database

var connection string //package level variable untuk menyimpan status koneksi database

func init() { //init function akan dijalankan otomatis saat package diimport
	connection = "Database connection POSTGRESQL established" //menginisialisasi variabel connection saat package diimport
}

func GetConnection() string { //fungsi untuk mendapatkan status koneksi database
	return connection //mengembalikan nilai variabel connection
}
