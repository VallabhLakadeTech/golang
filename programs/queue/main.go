package main

import "fmt"

func main() {

	myQueue := Queue{
		size: 5,
		data: make([]int, 5),
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

	fmt.Println("--------------Circular Queue-----------")
	createcircularqueue()
}
