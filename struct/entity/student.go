package entity

type Student struct { //struct untuk merepresentasikan data mahasiswa
	Name, Address string //field Name dan Address bertipe string
	Age           int    //field Age bertipe int
	Married       bool   //field Married bertipe bool
}

func (student Student) Person(name string) { //method Person dengan parameter name bertipe string
	println("hello", name, "my name is", student.Name) //menampilkan sapaan dengan nama parameter dan nama mahasiswa
}
func (student Student) Info() { //method Info untuk menampilkan informasi mahasiswa tanpa parameter
	println("Nama:", student.Name)       //menampilkan nama
	println("Umur:", student.Age)        //menampilkan umur
	println("Alamat:", student.Address)  //menampilkan alamat
	println("Menikah:", student.Married) //menampilkan status menikah

}
