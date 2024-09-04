package topics

import "fmt"

func Channels() {

	// ch := make(chan string, 5)
	// go printSomething("I am being sent through a channel", ch)

	// for data := range ch {
	// 	fmt.Println(data)
	// }

	// Buffered and unbuffered channels
	deadlock()

}

func printSomething(str string, ch chan string) {
	defer close(ch)
	for i := 0; i < 5; i++ {
		//processing
		ch <- str
	}
	// Dont close this channel and see what happens
}

func deadlock() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	data := <-ch1
	fmt.Println("data in deadlock :: ", data)
	data = <-ch1
	fmt.Println("data in deadlock :: ", data)
	close(ch1)
	// Use unbuffered channel to get rid of deadlock
}
