package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// usingOneChan()
	usingTwoChan()
}

func usingOneChan() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	printOdd := func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			<-ch
			fmt.Println("Odd:", i)
			ch <- 1
		}
	}

	printEven := func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-ch
			fmt.Println("Even: ", i)
			if i != 10 {
				ch <- 1
			}
		}
	}

	go printEven(ch, &wg)
	go printOdd(ch, &wg)
	ch <- 1
	wg.Wait()
}

func usingTwoChan() {

	oddCh := make(chan int)
	evenCh := make(chan int)

	go func() {
		defer close(oddCh)
		defer close(evenCh)
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				evenCh <- i
			} else {
				oddCh <- i
			}
		}
	}()

	for oddCh != nil || evenCh != nil {
		time.Sleep(20 * time.Millisecond)
		select {
		case odd, ok := <-oddCh:
			if ok {
				fmt.Println("Odd: ", odd)
			} else {
				oddCh = nil
			}

		case even, ok := <-evenCh:
			if ok {
				fmt.Println("Odd: ", even)
			} else {
				evenCh = nil
			}
		}
	}

}
