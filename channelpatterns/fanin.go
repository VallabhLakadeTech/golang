package main

import "fmt"

// Taking output of multiple channels and accumulating into one

func producer(output chan<- int, incrementer int) {

	for i := 0; i < 3; i++ {
		output <- i + incrementer
	}
	close(output)
}

func aggregator(producer1, producer2 <-chan int, output chan<- int) {
	defer close(output)
	for producer1 != nil || producer2 != nil {
		select {
		case value1, open1 := <-producer1:
			if open1 {
				output <- value1
			} else {
				producer1 = nil
			}
		case value2, open2 := <-producer2:
			if open2 {
				output <- value2
			} else {
				producer2 = nil
			}
		}
	}

}

func fanin() {

	output := make(chan int)
	producer1 := make(chan int)
	producer2 := make(chan int)

	go producer(producer1, 0)
	go producer(producer2, 10)
	go aggregator(producer1, producer2, output)

	for result := range output {
		fmt.Println(result)
	}

}
