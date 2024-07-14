package main

import (
	"fmt"
	"sync"
)

var (
	wg    sync.WaitGroup
	mu    sync.Mutex
	count int
)

func increment() {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	count = count + 1

}

func mutexlock() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go increment()
	}
	wg.Wait()
	fmt.Println("Count: ", count)
}
