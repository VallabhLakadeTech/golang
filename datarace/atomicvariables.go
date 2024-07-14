package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32

func atomicIncrement() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				atomic.AddInt32(&counter, 1)
			}
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Total counter value: ", counter)
}
