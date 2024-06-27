package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(num int) {
			fmt.Println(factorial(num + 1))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func factorial(num int) int {
	if num < 2 {
		return num
	}
	return num * factorial(num-1)
}
