package main

import "fmt"

func mapping(slice []string) map[string]int {
	banyak := make(map[string]int)

	for _, value := range slice {
		banyak[value]++
	}

	return banyak
}

func main() {
	fmt.Println(mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) //map [adi: 1 asd: 2 qwe: 3]
	fmt.Println(mapping([]string{}))
}
