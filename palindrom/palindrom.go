package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan kata: ")
	input, _ := reader.ReadString('\n')

	// bersihkan newline
	input = strings.TrimSpace(input)

	if isPalindrome(input) {
		fmt.Println("Output: true → kata ini palindrom")
		fmt.Println("Penjelasan: kata ini termasuk palindrom, karena kalau dibalik output nya tetap sama dengan kata yang dimasukkan.")
	} else {
		fmt.Println("Output: false → kata ini bukan palindrom")
		fmt.Println("Penjelasan: kata ini tidak termasuk palindrom karena kalau dibalik menghasilkan ", input)
	}
}
