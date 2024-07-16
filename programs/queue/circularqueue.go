package main

import "fmt"

type circularqueue struct {
	size   int
	isFull bool
	data   []int
	front  int
	rear   int
}

func (queue *circularqueue) addToRear(num int) {

	if queue.isFull && queue.front == queue.rear {
		fmt.Println("Queue is full")
		return
	}
	queue.data[queue.rear] = num
	queue.rear = (queue.rear + 1) % queue.size
	if (queue.rear + 1) == queue.size {
		queue.isFull = true
	}
	fmt.Println("Add number ", num)
}

func (queue *circularqueue) removeFromFront() {
	if queue.front == queue.rear && !queue.isFull {
		fmt.Println("Queue is empty")
		return
	}
	num := queue.data[queue.front]
	queue.front = (queue.front + 1) % queue.size
	fmt.Println("Removed ", num)
}

func createcircularqueue() {

	myQueue := circularqueue{
		data: make([]int, 5),
		size: 5,
	}

	myQueue.addToRear(3)
	myQueue.addToRear(5)
	myQueue.addToRear(8)
	myQueue.removeFromFront()
	myQueue.addToRear(2)
	myQueue.addToRear(7)
	myQueue.addToRear(6)
	myQueue.addToRear(9)
	myQueue.removeFromFront()

}
