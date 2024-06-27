package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)
	stopChannel := make(chan bool)

	go printSomething(channel1, channel2, stopChannel)

	for {
		select {
		case <-stopChannel:
			close(channel1)
			close(channel2)
			close(stopChannel)
			return
		case msg1 := <-channel1:
			fmt.Println("Message from channel1: ", msg1)
		case msg2 := <-channel2:
			fmt.Println("Message from channel2: ", msg2)

		}
	}
}

func printSomething(channel1, channel2 chan string, stopChannel chan bool) {

	timer500ms := time.NewTicker(500 * time.Millisecond)
	timer300ms := time.NewTicker(300 * time.Millisecond)
	timeout := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-timeout.C:
			timer500ms.Stop()
			timer300ms.Stop()
			timeout.Stop()
			stopChannel <- true
		case <-timer500ms.C:
			channel1 <- "after 500ms"
		case <-timer300ms.C:
			channel2 <- "after 300ms"
		}

	}

}
