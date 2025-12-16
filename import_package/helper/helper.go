package helper

type class struct {
	Name string
	Age  int
	NoID int
	// id   int //unexported field karena diawali dengan huruf kecil
}

func GetClass(name string, age int, noID int) *class {
	return &class{
		Name: name,
		Age:  age,
		NoID: noID,
		// id:   id, //unexported field tidak bisa diakses dari package lain jika ingin diakses harus melalui method
	}
}
