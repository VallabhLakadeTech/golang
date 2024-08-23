package main

import (
	"fmt"
	"time"
)

func main() {
	pingpong()
	buffered()
}

func pingpong() {

	ch := make(chan string)
	stop := true
	go func() {
		for stop {
			msg := <-ch
			fmt.Println(msg)
			time.Sleep(50 * time.Millisecond)
			ch <- "ping"
		}
	}()

	go func() {
		for stop {
			msg := <-ch
			fmt.Println(msg)
			time.Sleep(50 * time.Millisecond)
			ch <- "pong"
		}
	}()

	ch <- "ping"
	time.Sleep(1 * time.Second)
	stop = true
	close(ch)
}

func buffered() {

	bufCh := make(chan string, 2)
	bufCh <- "Hi"
	fmt.Println(<-bufCh)
	bufCh <- "Vallabh"
	fmt.Println(<-bufCh)
	close(bufCh)
}
