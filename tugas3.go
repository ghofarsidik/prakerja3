package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	banyak := make(map[rune]int)

	// Menghitung berapa kali setiap digit muncul dalam string
	for _, code := range angka {
		banyak[code]++
	}

	pilih1 := []int{}
	for _, code := range angka {
		if banyak[code] == 1 {
			pilih1 = append(pilih1, int(code-'0'))
		}
	}

	return pilih1
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
