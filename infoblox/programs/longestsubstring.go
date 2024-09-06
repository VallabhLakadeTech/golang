package programs

import "fmt"

/*
"Given a string, find the length of the longest substring without repeating characters.

Examples:

Given ""abcabcbb"", the answer is ""abc"", which the length is 3.

Given ""bbbbb"", the answer is ""b"", with the length of 1.

Given ""pwwkew"", the answer is ""wke"", with the length of 3.
Note that the answer must be a substring, ""pwke"" is a subsequence and not a substring."
*/
func LongestSubstring() {

	inputString := "bbbbb"

	longestSubstring := 0
	start, end := 0, 0
	charSet := make(map[rune]int)
	for index, char := range inputString {
		if _, ok := charSet[char]; !ok {
			charSet[char] = 1
			end++
		} else {
			difference := end - start
			if start == end {
				difference = 1
			}
			longestSubstring = max(difference, longestSubstring)
			start = index
			end = start + 1
			charSet = make(map[rune]int)
			charSet[char] = 1

		}
	}
	fmt.Println("Length of longest substring: ", longestSubstring)
}
