package main

import (
	"fmt"
	"math"
)

func main() {
	//package math
	// math round untuk pembulatan angka
	fmt.Println("Round 3.7:", math.Round(3.7))
	fmt.Println("Round 3.2:", math.Round(3.2))

	// math.MaxFloat32 bisa jika math 64 maka math.MaxFloat64
	fmt.Println("Max Float32:", math.MaxFloat32)
	fmt.Println("Min Int64:", math.MinInt64)

	// package math digunakan untuk operasi matematika
	fmt.Println("Square root of 16 is:", math.Sqrt(16))
	fmt.Println("Value of Pi is:", math.Pi)

	// contoh penggunaan fungsi Exp dari package math
	fmt.Println("Exponential of 1 is:", math.Exp(1))

	// contoh penggunaan fungsi Pow dari package math
	fmt.Println("2 raised to the power of 3 is:", math.Pow(2, 3))

	// contoh penggunaan fungsi Sin dari package math
	fmt.Println("Sine of Pi/2 is:", math.Sin(math.Pi/2))

	// contoh penggunaan fungsi Log dari package math
	fmt.Println("Natural logarithm of E is:", math.Log(math.E))

}
