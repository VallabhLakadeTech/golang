package main

import (
	"fmt"
	"time"
)

// Goroutine will wait for data and parent will send the task
func waitForTask() {

	ch := make(chan string)

	go func() {

		d := <-ch
		fmt.Println("Child got data ", d)

	}()
	time.Sleep(time.Millisecond * 4)
	fmt.Println("Parent sending signal")
	ch <- "vallabh"
	time.Sleep(time.Millisecond * 2)
	close(ch)
}
