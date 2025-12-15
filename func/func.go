package main

import (
	"fmt"
	"strings"
)

// struct untuk menyimpan gabungan nama & umur
type User struct {
	Name string
	Age  int
}

type Verify func(string) bool // deklarasi tipe fungsi Verify

// Fungsi tanpa parameter dan tanpa return
func PrintName() {
	fmt.Println("Nama saya Rizqi Fauzi")
}

// Fungsi tanpa parameter dan dengan return
func listNames() []string {
	names := []string{"Rizqi", "Fauzi", "Ryco", "Dwi", "Aji"}
	for i, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", i+1, name)
	}
	return names
}

func average(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	total := 0
	for _, num := range numbers {
		total += num
	}
	return float64(total) / float64(len(numbers))
}

// Returning Multiple Values → menggabungkan dua slice ke slice User
func livinguser(names []string, ages []int) []User {
	var users []User
	for i := 0; i < len(names) && i < len(ages); i++ {
		users = append(users, User{
			Name: names[i],
			Age:  ages[i],
		})
	}
	return users
}

// anonymous function
func verifyName(name string, verify Verify) {
	if verify(name) {
		fmt.Println("Verifikasi berhasil", name)
	} else {
		fmt.Println("Verifikasi gagal", name)
	}
}

func getfullname() (firstname string, middlename string, lastname string) {
	firstname = "Rizqi"
	middlename = "Fauzi"
	lastname = "Putra"
	return
}

func factorialages(n int) int { // ✔️ fungsi factorial non recursive yang sudah benar
	result := 1               // inisialisasi result dengan 1
	for i := 1; i <= n; i++ { // loop dari 1 sampai n
		result *= i // kalikan result dengan i
	}
	return result // kembalikan result
}

// ✔️ fungsi recursive factorial yang sudah benar
func factorialRecursive(n int) int { // fungsi factorial recursive
	if n == 1 { // basis kasus
		return 1 // jika n sama dengan 1, kembalikan 1
	} else { // kasus rekursif
		return n * factorialRecursive(n-1) // kalikan n dengan hasil factorialRecursive(n-1)
	}
}

func main() {
	PrintName()

	fmt.Println("\nMenggunakan range untuk mengambil slice:")
	names := listNames()
	fmt.Println("\nSemua nama sebagai slice:", strings.Join(names, ", "))

	ages := []int{17, 20, 23, 34, 55}
	fmt.Printf("Rata-rata umur: %.2f\n", average(ages))

	users := livinguser(names, ages)
	fmt.Println("\nData nama + umur:")
	for _, u := range users {
		fmt.Printf("Nama: %s → Umur: %d\n", u.Name, u.Age)
	}

	// anonymous function verify
	verify := func(name string) bool {
		name = strings.ReplaceAll(name, " ", "")
		return len(name) > 3
	}

	for _, name := range names {
		verifyName(name, verify)
	}

	// multiple return
	a, b, c := getfullname()
	fullname := fmt.Sprintf("%s %s %s", a, b, c)
	fmt.Println("Fullname:", fullname)

	loop := factorialages(5)                                 // non recursive
	fmt.Println("Factorial 5 adalah (non Recursive):", loop) // non recursive

	recursion := factorialRecursive(1)                        // recursive
	fmt.Println("Factorial 5 adalah (Recursive):", recursion) // recursive
}
