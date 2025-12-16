package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct { // mendefinisikan struct dengan tag
	Name string `required:"true" max:"10" min:"1"` // tag untuk field Name dengan nilai "true" dan "10"
	age  int    `min:"0"`                          // tag untuk field age dengan nilai "0"
}

type User struct { // mendefinisikan struct tidak ada tag
	Name string
	Age  int
}

// mendefinisikan fungsi ValidateStruct untuk melakukan validasi struct
// dengan menggunakan tag required dari validate struct dan menggunakan reflect
// keluaran fungsi adalah boolean
func ValidateStruct(data interface{}) bool {
	tag := reflect.TypeOf(data) // mendapatkan tipe data dari data
	// lakukan validasi berdasarkan tag
	for i := 0; i < tag.NumField(); i++ { // iterasi melalui setiap field
		field := tag.Field(i)                    // mendapatkan informasi field
		if field.Tag.Get("required") == "true" { // jika tag "required" bernilai "true"
			// lakukan validasi required
			value := reflect.ValueOf(data).Field(i) // mendapatkan nilai field
			if value.IsZero() {                     // jika nilai field kosong false
				return false
			}

		}
	}
	return true
}

func main() {
	sample := MyStruct{ // membuat instance dari MyStruct
		Name: "Hello",
		age:  42,
	}
	// mendapatkan reflect.Value dari instance sample
	var samplevalue reflect.Value = reflect.ValueOf(sample)
	// menggunakan reflect untuk mendapatkan informasi tentang struct sample
	fmt.Println("Number of fields", samplevalue.NumField())        // mencetak jumlah field dalam struct
	fmt.Println("Field 0 name:", samplevalue.Type().Field(0).Name) // mencetak nama field pertama
	fmt.Println("Field 1 name:", samplevalue.Type().Field(1).Name) // mencetak nama field kedua
	// mendapatkan tag dari field pertama
	fmt.Println("Field 0 tag required:", samplevalue.Type().Field(0).Tag.Get("required")) // mencetak nilai tag "required"
	fmt.Println("Field 0 tag max:", samplevalue.Type().Field(0).Tag.Get("max"))           // mencetak nilai tag "max"
	fmt.Println("Field 0 tag min:", samplevalue.Type().Field(1).Tag.Get("min"))           // mencetak nilai tag "min" (tidak ada, jadi kosong)

	// melakukan validasi struct menggunakan fungsi ValidateStruct
	isValid := ValidateStruct(sample)
	fmt.Println("Is Mystruct valid?", isValid) // mencetak hasil validasi

	// melakukan validasi struct menggunakan fungsi ValidateStruct tetapi tidak ada tag field
	user := User{"", 0}
	fmt.Println("Is User valid?", ValidateStruct(user))

}
