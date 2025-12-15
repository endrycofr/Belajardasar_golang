package main

import ( //import package entity dan service
	"fmt"
	"repo/entity"
	"repo/service"
)

func main() {
	user := entity.User{ //membuat instance struct User dari package entity
		Name:  "Rizqi Fauzi",
		Email: "rizqi@email.com",
		Age:   20,
	}
	userService := service.NewUserService(user) //membuat instance struct UserService dari package service dengan dependency injection seperti (repo, service, dll)

	// ================= SHOW USER =================
	//  Mengirim data user ke service
	//  Service bertugas mengolah & menampilkan data, bukan main

	userService.ShowUser() //memanggil method ShowUser dari struct UserService

	// ================= REPO FUNCTION =================
	// 	Mengirim string "good"
	//  Service memproses data lewat repository
	//  Hasil dikembalikan ke main
	//  Ditampilkan ke console

	data := userService.RepoFunction("good") //memanggil method RepoFunction dari struct UserService
	fmt.Println("Result:", data)             //menampilkan hasil dari method RepoFunction

	// ================= STATUS SERVICE =================
	// Membuat instance struct UserService dari package service tanpa dependency injection seperti (repo, service, dll)
	// 	Mengirim data user ke service
	//  Service bertugas mengolah & menampilkan data, bukan main
	//  Service memproses data lewat repository
	//  Hasil dikembalikan ke main
	//  Ditampilkan ke console

	statusService := service.StatusUserService() //membuat instance struct UserService dari package service tanpa dependency injection

	status, err := statusService.SetRepository(0) //memanggil method SetRepository dari struct UserService
	if err != nil {
		fmt.Println("error:", err) //menampilkan error jika ada
		return                     //menghentikan eksekusi program jika ada error
	}

	fmt.Println(status) //menampilkan hasil dari method SetRepository
}
