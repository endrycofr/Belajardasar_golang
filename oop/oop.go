package main

import "fmt"

// Struct sebagai "class"
type Hewan struct { // mendefinisikan struct Hewan
	Nama  string
	Suara string
}

// Method pada struct (Encapsulation)
func (h Hewan) Berbicara() string { // method untuk struct Hewan
	return fmt.Sprintf("%s bersuara:  %s", h.Nama, h.Suara) // mengembalikan format string dengan nama dan suara hewan
}

// Embedding sebagai "inheritance"
type Kucing struct { // mendefinisikan struct Kucing
	Hewan // mewarisi struct Hewan
}

// Polymorphism melalui interface
type Speaker interface { // mendefinisikan interface Speaker
	Berbicara() string // method yang harus diimplementasikan oleh struct yang mengimplementasikan interface ini
}

func buatBersuara(s Speaker) { // fungsi yang menerima parameter bertipe Speaker
	fmt.Println(s.Berbicara()) // mencetak suara dari objek yang mengimplementasikan interface Speaker
}

func main() {
	kucing := Kucing{
		Hewan{"Kitty", "Meong"},
	}

	buatBersuara(kucing)
}
