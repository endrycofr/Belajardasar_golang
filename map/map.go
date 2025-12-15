package main

func main() {
	type name struct { // deklarasi tipe baru name sebagai struct
		firstName string // deklarasi field firstName dengan tipe data string
		lastName  string // deklarasi field lastName dengan tipe data string
		age       int    // deklarasi field age dengan tipe data int
		weight    int    // deklarasi field weight dengan tipe data int
	}

	person := make(map[string]name) // deklarasi variabel person sebagai map dengan key bertipe string dan value bertipe name

	person["adit"] = name{ // mengisi map dengan key "adit"
		firstName: "Adit",    // mengisi field firstName dengan nilai "Adit"
		lastName:  "Santoso", // mengisi field lastName dengan nilai "Santoso"
		age:       28,        // mengisi field age dengan nilai 28
		weight:    75,        // mengisi field weight dengan nilai 75
	}
	person["bryan"] = name{ // mengisi map dengan key "bryan"
		firstName: "Bryan", // mengisi field firstName dengan nilai "Bryan"
		lastName:  "Park",  // mengisi field lastName dengan nilai "Park"
		age:       30,      // mengisi field age dengan nilai 30
		weight:    80,      // mengisi field weight dengan nilai 80
	}

	println(person["adit"].firstName + " " + person["adit"].lastName) // mencetak nilai field firstName dan lastName dari key "adit"

	println(person["adit"].age)    // mencetak nilai field age dari key "adit"
	println(person["adit"].weight) // mencetak nilai field weight dari key "adit"

	println(person["bryan"].firstName + " " + person["bryan"].lastName) // mencetak nilai field firstName dan lastName dari key "adit"

	println("sebelum di delete", len(person)) // mencetak panjang dari map person

	delete(person, "bryan")                   // menghapus key "bryan" dari map person
	println("sesudah di delete", len(person)) // mencetak panjang dari map person
}
