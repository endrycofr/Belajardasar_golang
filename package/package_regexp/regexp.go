package main

import (
	"fmt"
	"regexp"
)

type login struct { // struct login berisi email dan password
	email    string
	password string
}

func main() {
	login := login{ // membuat instance struct login
		email:    "espadalois@example.com",
		password: "password123",
	}

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$` // pattern email
	matched, _ := regexp.MatchString(pattern, login.email)        // melakukan pencocokan pattern dengan email

	if matched { // kondisi jika email valid dan tidak valid
		println("Email valid")
	} else {
		println("Email tidak valid")
	}
	passwordPattern := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$` // pattern password
	matched, _ = regexp.MatchString(passwordPattern, login.password)    // melakukan pencocokan pattern dengan password
	if matched {
		println("Password valid")
	} else {
		println("Password tidak valid")
	}

	// SEARCH STRING DENGAN REGEXP
	re := regexp.MustCompile(`espada lois`)                                         // membuat regexp untuk mencari kata "espada lois"
	fmt.Println(re.FindAllString("nama saya espada lois pada kaka espada lois", 1)) // mencari semua kemunculan kata "espada lois" dalam string
}
