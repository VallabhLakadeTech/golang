package programs

import (
	"fmt"
)

func ReverseString() {

	inputString := "álphãbét"
	myRune := []rune(inputString)

	for i, j := 0, len(myRune)-1; i < j; i, j = i+1, j-1 {
		myRune[i], myRune[j] = myRune[j], myRune[i]
	}
	fmt.Println("Reverse string", string(myRune))
}
