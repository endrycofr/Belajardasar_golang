package main

import (
	"container/list"
	"fmt"
)

func main() {
	value := list.New()
	// menambahkan elemen ke dalam list
	// menggunakan PushBack (dari belakang) dan PushFront (dari depan)
	value.PushBack(10)
	value.PushBack(20)
	value.PushFront(5)
	value.PushFront(1)

	// iterasi elemen dalam list dari depan ke belakang
	for element := value.Front(); element != nil; element = element.Next() {
		fmt.Println("Element value from list front:", element.Value)
	}

	// iterasi elemen dalam list dari belakang ke depan
	for element := value.Back(); element != nil; element = element.Prev() {
		fmt.Println("Element value from list back:", element.Value)
	}
	// mengetahui panjang list
	fmt.Println("Length of list:", value.Len())
	// mengakses elemen depan dan belakang list
	fmt.Println("Front element:", value.Front().Value)
	// mengakses elemen belakang list
	fmt.Println("Back element:", value.Back().Value)
}
