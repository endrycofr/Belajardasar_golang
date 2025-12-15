package main

import "fmt"

type Country struct { //struct Country
	city     string
	region   string
	province string
	zip      int
}

func changeCity(country *Country) { //fungsi changeCity yang menerima pointer ke Country struct
	country.city = "Bandung" //mengubah nilai city menjadi "Bandung"
}

func (man *Country) Man() { //method countryman yang menerima pointer ke Country struct
	man.province = "Jawa Barat" //mengubah nilai province menjadi "Jawa Barat"
}

func main() {
	myCountry := Country{ //membuat instance dari Country struct di memory stack
		city:     "Jakarta",
		region:   "Jabodetabek",
		province: "DKI Jakarta",
		zip:      12345,
	}
	var countries *Country = &myCountry
	//membuat pointer ke myCountry struct diatas menggunakan & operator untuk mendapatkan alamat memorynya dan disimpan dalam variabel countries
	changeCity(countries)
	fmt.Println(myCountry)

	//===== pointer method =========
	countries.Man()                                   //memanggil method Man menggunakan pointer countries
	fmt.Println("My province is", countries.province) //menampilkan nilai province setelah diubah oleh method Man
	//output: My province is Jawa Barat ,jadi jika print myCountry maka province nya sudah berubah menjadi Jawa Barat

}
