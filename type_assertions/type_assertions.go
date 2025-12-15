package main

func checktype(name interface{}) { //fungsi untuk memeriksa tipe data dari parameter name yang bertipe interface{}
	switch value := name.(type) { //melakukan type assertion menggunakan switch
	case string: //jika tipe data adalah string
		println("Name is a string:", value)
	case int: //jika tipe data adalah int
		println("Name is an integer:", value)
	default: //jika tipe data tidak dikenali
		println("Unknown type")
	}
}

func main() {
	checktype("Alice") //memanggil fungsi checktype dengan parameter string
	checktype(42)      //memanggil fungsi checktype dengan parameter int
	checktype(true)    //memanggil fungsi checktype dengan parameter boolean
}
