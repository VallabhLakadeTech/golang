package main

import (
	"fmt"
	"sort"
)

/*
Problem statement: Find all possible pair whose addition results in the given sum
*/
func main() {

	inputArr := []int{2, 7, 1, 8, 3, 6, 5}
	// 1,2,3,5,6,7,8
	sum := 10

	sort.Ints(inputArr)

	i := 0
	j := len(inputArr) - 1
	fmt.Println(inputArr)
	for i < j {
		addition := inputArr[i] + inputArr[j]
		if addition < sum {
			i = i + 1
			continue
		} else if addition > sum {
			j = j - 1
			continue
		}
		fmt.Printf("Pair found: [%v,%v]\n", inputArr[i], inputArr[j])
		i = i + 1
		j = j - 1
	}

}
