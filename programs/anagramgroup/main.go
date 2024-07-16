package main

import (
	"fmt"
	"strings"
)

func main() {
	stringList := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	anagramList := groupanagram(stringList)
	// Print grouped anagrams
	for _, group := range anagramList {
		fmt.Println(strings.Join(group, ", "))
	}
}

func frequency(str string) [26]int {

	var freq [26]int
	for _, ch := range str {
		freq[ch-'a']++
	}
	return freq
}

func groupanagram(stringList []string) map[[26]int][]string {
	// make(map[string]int)
	anagrams := make(map[[26]int][]string)

	for _, str := range stringList {
		key := frequency(str)
		anagrams[key] = append(anagrams[key], str)
	}
	return anagrams
}
