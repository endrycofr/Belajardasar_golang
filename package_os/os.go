package main

import packageos "package_os/package_os"

func main() {
	packageos.PrintArgs()               //memanggil fungsi untuk mencetak argumen dari command line
	hostname := packageos.GetHostname() //memanggil fungsi untuk mendapatkan hostname sistem operasi
	println("Hostname:", hostname)      //menampilkan hostname

}
