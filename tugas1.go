package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	tanda := make(map[string]bool)

	for _, nama := range arrayA {
		tanda[nama] = true
	}

	for _, nama := range arrayB {
		tanda[nama] = true
	}

	gabung := make([]string, 0, len(tanda))
	for nama := range tanda {
		gabung = append(gabung, nama)
	}
	return gabung
}

func main() {
	//test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil", "akuma, "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"}

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil", "jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}
