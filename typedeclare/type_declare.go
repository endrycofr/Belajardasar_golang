package main

func main() {
	type ( // deklarasi tipe baru menggunakan kata kunci type atau alias
		nomorKTP     string // deklarasi tipe baru bernama nomorKTP dengan tipe dasar string
		namaDepan    string // deklarasi tipe baru bernama namaDepan dengan tipe dasar string
		namaBelakang string // deklarasi tipe baru bernama namaBelakang dengan tipe dasar string
	)
	var nama namaDepan = "Budi"            // deklarasi variabel dengan tipe data string namaDepan
	var namaAkhir namaBelakang = "Santoso" // deklarasi variabel dengan tipe data string namaBelakang
	println("Nama lengkap saya adalah:", nama, namaAkhir)

	var ktp nomorKTP = "1234567890123456"  // deklarasi variabel dengan tipe data string nomorKTP
	println("Nomor KTP saya adalah:", ktp) //

	// contoh lain penggunaan tipe data baru

	var ktpSaya nomorKTP = "1234567890123456"
	println("Nomor KTP saya adalah:", ktpSaya)

}
