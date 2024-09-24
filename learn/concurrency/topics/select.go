package topics

import (
	"fmt"
	"time"
)

func Select() {
	SelectStatement()
}

func SelectStatement() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for i := 0; i < 15; i++ {
			time.Sleep(time.Millisecond * 75)
			ch1 <- "first go routine"
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 15; i++ {
			time.Sleep(time.Millisecond * 150)
			ch2 <- "second go routine"
		}
		close(ch2)
	}()

	myTicker := time.NewTicker(time.Second * 1)

	// for {
	for ch1 != nil || ch2 != nil {
		select {
		case data, open := <-ch1:
			if open {
				fmt.Println(data)
			} else {
				ch1 = nil
			}
		case data, open := <-ch2:
			if open {
				fmt.Println(data)
			} else {
				ch2 = nil
			}
		case <-myTicker.C:
			fmt.Println("I will run after every 1 second")
		}
	}

}
