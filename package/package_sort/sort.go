package main

import (
	"fmt"
	"sort"
)

type user struct { // mendefinisikan struct user dengan field name dan age
	name string
	age  int
}
type sortByName []user // mendefinisikan tipe baru sortByName yang merupakan slice dari user

func (a sortByName) Len() int { // mengimplementasikan method Len untuk interface sort.Interface
	return len(a)
}
func (a sortByName) Swap(i, j int) { // mengimplementasikan method Swap untuk interface sort.Interface
	a[i], a[j] = a[j], a[i]
}
func (a sortByName) Less(i, j int) bool { // mengimplementasikan method Less untuk interface sort.Interface
	return a[i].age < a[j].age
}

func main() {
	users := []user{ // membuat slice dari user
		{name: "Alice", age: 30},
		{name: "Bob", age: 25},
		{name: "Charlie", age: 35},
	}

	sort.Sort(sortByName(users)) // mengurutkan slice users berdasarkan age menggunakan sort.Sort

	fmt.Println("Users sorted by age:") // mencetak header
	for _, u := range users {           // mencetak setiap user dalam slice yang sudah diurutkan
		fmt.Printf("Name: %s, Age: %d\n", u.name, u.age)
	}

}
