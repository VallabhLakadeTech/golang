package programs

import (
	"fmt"
	"sync"
)

func NaturalNumbers(num int) {

	// Unordered sequence of numbers
	var wg sync.WaitGroup

	for i := 1; i <= num; i++ {
		wg.Add(1)
		go func(currentNum int) {
			fmt.Println(currentNum)
			defer wg.Done()
		}(i)
	}
	wg.Wait()

}
