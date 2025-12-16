package main

import (
	"fmt"
	"strconv"
)

// ConvTypeData struct untuk menyimpan hasil konversi berbagai tipe data
type ConvTypeData struct {
	IntToString     string
	StringToInt     int
	FloatToString   string
	StringToFloat   float64
	BooleanToString string
	StringToBoolean bool
}

func main() {
	data := ConvTypeData{}

	// int -> string
	data.IntToString = strconv.Itoa(123)

	// string -> int
	data.StringToInt, _ = strconv.Atoi("123")

	// float -> string
	data.FloatToString = strconv.FormatFloat(3.14, 'f', -1, 64)

	// string -> float
	data.StringToFloat, _ = strconv.ParseFloat("3.14", 64)

	// bool -> string
	data.BooleanToString = strconv.FormatBool(true)

	// string -> bool
	data.StringToBoolean, _ = strconv.ParseBool("true")

	// menampilkan hasil konversi beserta tipe data
	fmt.Printf("Type IntToString: %T\n | value: %v\n", data.IntToString, data.IntToString)
	fmt.Printf("Type StringToInt: %T\n | value: %v\n", data.StringToInt, data.StringToInt)
	fmt.Printf("Type FloatToString: %T\n | value: %v\n", data.FloatToString, data.FloatToString)
	fmt.Printf("Type StringToFloat: %T\n | value: %v\n", data.StringToFloat, data.StringToFloat)
	fmt.Printf("Type BooleanToString: %T\n | value: %v\n", data.BooleanToString, data.BooleanToString)
	fmt.Printf("Type StringToBoolean: %T\n | value: %v\n", data.StringToBoolean, data.StringToBoolean)

	// strconv.Atoi untuk konversi string ke integer dengan penanganan error
	value, err := strconv.Atoi("20000")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Value: %d\n | Type: %T", value, value)
	}
}
