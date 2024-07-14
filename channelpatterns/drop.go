package main

import (
	"fmt"
	"time"
)

// We have a buffered channel and we only process data till the buffer point and drop the others
func drop() {

	ch := make(chan int, 50)

	go func() {
		for data := range ch {
			fmt.Println("Child handling data ", data)
			time.Sleep(time.Millisecond * 2)
		}
	}()

	for i := 0; i < 1000; i++ {

		select {
		case ch <- i:
			time.Sleep(time.Millisecond * 1)
			fmt.Println("Data is being pushed from parent", i)
		default:
			fmt.Println("Data is being dropped from access request from parent", i)
		}

	}
	time.Sleep(time.Millisecond * 5)
	close(ch)
}
