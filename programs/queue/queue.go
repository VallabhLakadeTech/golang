package main

import "fmt"

type Queue struct {
	size        int
	front, rear int
	data        []int
}

func (queue *Queue) addToRear(num int) {

	if queue.size == queue.rear {
		fmt.Println("queue is full")
		return
	}
	queue.data[queue.rear] = num
	queue.rear++
	fmt.Println("Added ", num)
}

func (queue *Queue) removeFromFront() {

	if queue.front == queue.rear {
		fmt.Println("queue is empty")
		return
	}
	num := queue.data[queue.front]
	queue.front++
	queue.rear--
	fmt.Println("Removed ", num)

}
