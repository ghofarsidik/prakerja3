package main

import (
	"fmt"
)

func munculSekali(angka string) int {
	banyak := make(map[rune]int)

	//
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
	fmt.Println(munculsekali("1234123"))
	fmt.Println(munculsekali("76523752"))
	fmt.Println(munculsekali("12345"))
	fmt.Println(munculsekali("1122334455"))
	fmt.Println(munculsekali("0872504"))
}
