package main

import (
	"fmt"
	"runtime"
	"time"
)

// We set runtime.GOMAXPROCS and run multiple gorutines which pool on the same data channel for inputs
func pooling() {

	ch := make(chan int)
	go pool(ch)

	for i := 0; i < 50; i++ {
		ch <- i
	}
	close(ch)

}

func pool(ch chan int) {

	runtime.GOMAXPROCS(2)
	for i := 0; i < 2; i++ {

		go func(process int) {
			for data := range ch {
				fmt.Printf("Data %v is being processed by process %v\n", data, process)
				time.Sleep(time.Millisecond * 4)
			}
		}(i)

	}

}
