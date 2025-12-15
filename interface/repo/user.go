package repo

import "fmt"

type UserRepo interface { //method signature kontrak interface
	GetName() string
	GetEmail() string
}

func NewUserRepo(val string) interface{} { //method signature kontrak interface kosong dan menerima string
	if val == "good" { //mengecek apakah val sama dengan "good"
		return true
	} else {
		return false
	}

}
func SetRepository(val interface{}) error { //method signature kontrak interface error dan menerima interface kosong
	if val != 1 { //mengecek apakah val tidak sama dengan 1
		return fmt.Errorf("invalid repository") //mengembalikan error jika val tidak sama dengan 1
	} else { //jika val sama dengan 1
		return nil //mengembalikan nil jika val sama dengan 1
	}

}
