package main

import (
	"fmt"
	"sync"
	"time"
)

type Fact struct {
	num       int64
	factorial int64
	worker    int
}

func main() {

	inputChannel := make(chan Fact, 10)
	outputChannel := make(chan Fact, 10)

	workerPoolSize := 4

	for i := 0; i < workerPoolSize; i++ {
		go work(inputChannel, outputChannel, i)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 1; i <= 20; i++ {
			fact := Fact{
				num: int64(i),
			}
			inputChannel <- fact
			time.Sleep(2 * time.Millisecond)
		}
		close(inputChannel)
	}()

	go func() {
		for i := 1; i <= 20; i++ {
			factResult := <-outputChannel
			fmt.Printf("Worker %v : Factorial of %v is %v\n", factResult.worker, factResult.num, factResult.factorial)
		}
		close(outputChannel)
		wg.Done()
	}()
	wg.Wait()
}

func work(inputChannel <-chan Fact, outputChannel chan<- Fact, worker int) {
	for input := range inputChannel {
		input.factorial = factorial(input.num)
		input.worker = worker
		outputChannel <- input
	}
}

func factorial(num int64) int64 {
	if num < 2 {
		return num
	}
	return num * factorial(num-1)
}
