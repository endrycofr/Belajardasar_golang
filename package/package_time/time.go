package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Package Time")
	//package time digunakan untuk mengelola waktu dan tanggal
	fmt.Println("Now Date Time:", now)
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())

	//membuat waktu tertentu dengan time.Date fungsi

	utc := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("UTC:", utc)

	//format waktu dengan Format fungsi time dan Parse fungsi time RFC3339
	timer := ("2006-01-02")
	parse, _ := time.Parse(timer, "2023-01-01")
	fmt.Println("Parse Date:", parse)
}
