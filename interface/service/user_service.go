package service

import (
	"fmt"
	"repo/repo"
)

type UserService struct { //dependency injection struct untuk menerima objek lain
	repo repo.UserRepo //menggunakan interface UserRepo
}

func NewUserService(r repo.UserRepo) *UserService { //constructor function untuk UserService struct
	return &UserService{repo: r} //mengembalikan pointer ke UserService dengan repo diisi dengan r
}
func StatusUserService() *UserService { //constructor function untuk UserService struct tanpa dependency injection
	return &UserService{} //mengembalikan pointer ke UserService dengan repo diisi dengan nil
}

func (s *UserService) ShowUser() { //method ShowUser untuk UserService struct
	fmt.Println("Name :", s.repo.GetName())  //memanggil method GetName dari interface UserRepo
	fmt.Println("Email:", s.repo.GetEmail()) //memanggil method GetEmail dari interface UserRepo
}

func (s *UserService) RepoFunction(val string) interface{} { //method RepoFunction untuk UserService struct
	return repo.NewUserRepo(val) //memanggil function NewUserRepo dari package repo
}

func (s *UserService) SetRepository(val interface{}) (string, error) { //method SetRepository menggunakan UserService struct yang tanpa dependency injection
	if err := repo.SetRepository(val); err != nil { //memanggil function SetRepository dari package repo
		return "", err //mengembalikan error jika ada
	}

	return "Repository berhasil diset", nil //mengembalikan string jika tidak ada error
}
