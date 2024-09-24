package main

import (
	"fmt"
)

func main() {
	fmt.Println("Is anagram: ", anagram("eat", "ate1"))

}

// eat,ate
func anagram(str1, str2 string) bool {

	anagramMap := make(map[rune]int)

	if len(str1) != len(str2) {
		return false
	}
	for key, value := range str1 {
		rune2 := rune(str2[key])
		if runeValue, ok := anagramMap[value]; !ok {
			anagramMap[value] = 1
		} else {
			anagramMap[value] = runeValue - 1
		}
		if runeValue, ok := anagramMap[rune2]; !ok {
			anagramMap[rune2] = 1
		} else {
			anagramMap[rune2] = runeValue - 1
		}
	}
	for _, value := range anagramMap {
		if value > 0 {
			return false
		}
	}
	return true
}
