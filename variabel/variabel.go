package main

import "fmt"

func main() {
	var name string = "rizqi"             // deklarasi variabel dengan tipe data string
	var age int = 20                      // deklarasi variabel dengan tipe data integer
	var height float32 = 170.5            // deklarasi variabel dengan tipe data float32
	var isStudent bool = true             // deklarasi variabel dengan tipe data boolean
	fmt.Println("Name:", name)            // mencetak nilai variabel name
	name = "dio"                          // mengubah nilai variabel name
	fmt.Println("Age:", age)              // mencetak nilai variabel age
	fmt.Println("Height:", height)        // mencetak nilai variabel height
	fmt.Println("Is Student:", isStudent) // mencetak nilai variabel isStudent

	fmt.Println("Name:", name) // mencetak nilai variabel name setelah diubah
	new_name := "dio"          // deklarasi variabel dengan tipe data string menggunakan short declaration
	fmt.Println("New Name:", new_name)
	new_age := 21 // deklarasi variabel dengan tipe data integer menggunakan short declaration
	fmt.Println("New Age:", new_age)
	fmt.Println("Height:", height)
	fmt.Println("Is Student:", isStudent)

	var (
		firstName string = "rizqi"
		lastName  string = "fauzi"
	)
	fmt.Println("Full Name:", firstName, lastName)

}
