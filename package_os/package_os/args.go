package packageos

import "os"

func PrintArgs() {
	args := os.Args[1:] //mengambil argumen dari command line kecuali nama program itu sendiri
	if len(args) == 0 { //jika tidak ada argumen yang diberikan
		println("No arguments provided.") //tampilkan pesan
		return                            //keluar dari program
	}
	for _, arg := range args { //iterasi melalui setiap argumen yang diberikan
		println(arg) //tampilkan argumen

	}
}
