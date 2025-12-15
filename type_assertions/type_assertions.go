package main

func checktype(name interface{}) {
	switch value := name.(type) {
	case string:
		println("Name is a string:", value)
	case int:
		println("Name is an integer:", value)
	default:
		println("Unknown type")
	}
}

func main() {
	checktype("Alice")
	checktype(42)
	checktype(true)
}
