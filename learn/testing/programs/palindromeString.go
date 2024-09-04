package programs

import (
	"fmt"
)

func PalindromeString(str string) {

	// if Reverse(str) {
	// 	fmt.Println("Its a palindrome")
	// 	return
	// }
	// fmt.Println("Its not a palindrome")

	if twoPointer(str) {
		fmt.Println("Its a palindrome")
		return
	}
	fmt.Println("Its not a palindrome")

}

// Not efficient program as string is immutable
func Reverse(str string) bool {
	var reverse string
	for _, v := range str {
		reverse = string(v) + reverse
	}
	return reverse == str
}

func twoPointer(str string) bool {
	runes := []rune(str)

	start := 0
	back := len(runes) - 1

	isPalindrome := true
	for start <= back {

		if runes[start] == runes[back] {
			start++
			back--
			continue
		}
		isPalindrome = false
		break
	}
	return isPalindrome
}
