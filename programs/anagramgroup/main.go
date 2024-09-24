package main

import (
	"fmt"
	"strings"
)

func main() {
	AnagramGroup()
}

func AnagramGroup() {
	stringList := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	anagramList := Groupanagram(stringList)
	// Print grouped anagrams
	for _, group := range anagramList {
		fmt.Println(strings.Join(group, ", "))
	}
}

func Frequency(str string) [26]int {

	var freq [26]int
	for _, ch := range str {
		freq[ch-'a']++
	}
	return freq
}

func Groupanagram(stringList []string) map[[26]int][]string {
	// make(map[string]int)
	anagrams := make(map[[26]int][]string)

	for _, str := range stringList {
		key := Frequency(str)
		anagrams[key] = append(anagrams[key], str)
	}
	return anagrams
}
