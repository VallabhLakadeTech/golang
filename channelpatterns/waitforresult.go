package main

import (
	"fmt"
	"time"
)

// Parent will wait for result to be received from child
func waitForResult() {

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Millisecond)
		ch <- "Vallabh"
		fmt.Println("Data inserted into channel")
	}()

	fmt.Println("Data received: ", <-ch)

}
