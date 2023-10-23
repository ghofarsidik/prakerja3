package main

import (
	"fmt"
)

func munculsekali(angka string) int {
	banyak := make(map[rune]int)

	for _, code := range angka {
		banyak[code]++
	}

	for _, code := range angka {
		if banyak[code] == 1 {
			return int(code - '0')
		}
	}
	return -1
}

func main() {
	fmt.Println(munculsekali("1234123"))
	fmt.Println(munculsekali("76523752"))
	fmt.Println(munculsekali("12345"))
	fmt.Println(munculsekali("1122334455"))
	fmt.Println(munculsekali("0872504"))
}
