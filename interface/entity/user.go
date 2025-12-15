package entity

type User struct { //struct User implements interface UserRepo
	Name  string
	Email string
	Age   int
}

func (u User) GetName() string { //mengimplementasi method GetName dari interface UserRepo
	return u.Name //mengembalikan nilai Name
}

func (u User) GetEmail() string { //mengimplementasi method GetEmail dari interface UserRepo
	return u.Email //mengembalikan nilai Email
}
