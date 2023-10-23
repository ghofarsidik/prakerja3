package main

import "fmt"

func munculSekali(input string) int {
	hitungan := make(map[rune]int)

	// Hitung berapa kali setiap karakter muncul dalam string
	for _, char := range input {
		hitungan[char]++
	}

	// Hitung berapa karakter yang hanya muncul sekali
	jumlahMunculSekali := 0
	for _, char := range input {
		if hitungan[char] == 1 {
			jumlahMunculSekali++
		}
	}

	return jumlahMunculSekali
}

func main() {
	fmt.Println(munculSekali("1234123")) // Output: 1 (karena hanya karakter "4" yang muncul sekali)
}
