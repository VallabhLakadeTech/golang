package programs

import (
	"fmt"
	"strings"
	"unicode"
)

func WordCount() {
	// Test paragraph
	paragraph := `Hello, world! This is a test paragraph.
	This paragraph is a test, and this is another test.`

	// Get word count
	wordCounts := findCount(paragraph)

	// Print word counts
	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func findCount(paragraph string) map[string]int {

	paragraph = strings.ToLower(paragraph)

	var stringBuilder strings.Builder

	for _, word := range paragraph {
		if unicode.IsLetter(word) || unicode.IsSpace(word) {
			stringBuilder.WriteRune(word)
		}
	}

	wordCount := make(map[string]int)
	words := strings.Fields(stringBuilder.String())
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}
