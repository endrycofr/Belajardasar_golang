package main

import "struct/entity" //import package entity

func main() { //main function
	student := entity.Student{ //membuat instance struct Student dari package entity
		Name:    "Rizqi Fauzi", //mengisi field Name
		Address: "Bandung",     //mengisi field Address
		Age:     20,            //mengisi field Age
		Married: false,         //mengisi field Married
	}

	student.Person("budi") //memanggil method Person dengan parameter
	student.Info()         //memanggil method Info dari instance student

}
