package main

import "fmt"

func main() {
	fmt.Println("Contoh penggunaan break dan continue dalam loop")
	for i := 1; i <= 10; i++ { // loop dari 1 hingga 10
		if i%2 == 0 { // jika i adalah angka genap
			fmt.Printf("Melompati angka genap: %d\n", i)
			continue // lewati angka genap
		}

		if i > 7 { // jika i lebih dari 7
			fmt.Printf("Berhenti pada angka lebih dari 7: %d\n", i) // tampilkan angka
			break                                                   // keluar dari loop jika i > 7
		}

		fmt.Printf("Angka: %d\n", i) // tampilkan angka ganjil
	}
}
