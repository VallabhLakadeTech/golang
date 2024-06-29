package main

import (
	"fmt"
	"time"
)

type Fact struct {
	num       int
	factorial int
	worker    int
}

func main() {

	inputChannel := make(chan Fact, 10)
	outputChannel := make(chan Fact, 10)

	workerPoolSize := 4

	for i := 0; i < workerPoolSize; i++ {
		go work(inputChannel, outputChannel, i)
	}

	for i := 1; i <= 10; i++ {
		fact := Fact{
			num: i,
		}
		inputChannel <- fact
		time.Sleep(2 * time.Millisecond)
	}
	close(inputChannel)

	for i := 1; i <= 10; i++ {
		factResult := <-outputChannel
		fmt.Printf("Worker %v : Factorial of %v is %v\n", factResult.worker, factResult.num, factResult.factorial)
	}
	close(outputChannel)

}

func work(inputChannel <-chan Fact, outputChannel chan<- Fact, worker int) {
	for input := range inputChannel {
		input.factorial = factorial(input.num)
		input.worker = worker
		outputChannel <- input

	}
}

func factorial(num int) int {
	if num < 2 {
		return num
	}
	return num * factorial(num-1)
}
