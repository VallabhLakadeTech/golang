package programs

import (
	"fmt"
	"sync"
)

func EvenOdd(num int) {

	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 2; i <= num; i = i + 2 {
			<-ch
			fmt.Println("Even: ", i)
			if i < num {
				ch <- i
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= num; i = i + 2 {
			<-ch
			fmt.Println("Odd: ", i)
			if i < num {
				ch <- i
			}
		}
	}()

	if num > 0 {
		ch <- 1
	}

	wg.Wait()

}
