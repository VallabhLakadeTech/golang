package main

import (
	"fmt"
	"sync"
)

var (
	count2 int
)

func increment2(counterCh chan int) {

	counterCh <- 1

}

func channel() {

	counterCh := make(chan int, 1000)
	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			increment2(counterCh)
		}()
	}
	go func() {
		for data := range counterCh {
			count2 = count2 + data
		}
	}()

	defer close(counterCh)
	wg.Wait()

	fmt.Println("Count through channel: ", count2)
}
