package main

import "fmt"

func work(tasks <-chan int, output chan<- int, workerID int) {

	for task := range tasks {
		fmt.Printf("Worker %v is working\n", workerID)
		output <- task * 5
	}

}

// Starting multiple goroutines to handle multiple tasks
// Advantage: better CPU utilization
func fanout() {

	workers := 3
	totalTasks := 15
	tasks := make(chan int, totalTasks)
	output := make(chan int, totalTasks)

	for i := 0; i < workers; i++ {
		go work(tasks, output, i)
	}
	for i := 0; i < totalTasks; i++ {
		tasks <- i
	}

	close(tasks)
	for i := 0; i < totalTasks; i++ {
		fmt.Println("Output: ", <-output)
	}

}
