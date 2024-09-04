package programs

import "fmt"

type AllowdQueueTypes interface {
	int | float64 | string
}

type QueueOperations[T AllowdQueueTypes] interface {
	enqueue(T)
	dequeue()
}

type Queue[T AllowdQueueTypes] struct {
	myQueue []T
	size    int
}

func (queue *Queue[T]) enqueue(data T) {
	if queue.size != len(queue.myQueue) {
		queue.myQueue = append(queue.myQueue, data)
		fmt.Println(queue.myQueue)
		return
	}
	fmt.Println("Queue is full")
}

func (queue *Queue[T]) dequeue() {

	if len(queue.myQueue) != 0 {
		fmt.Println("Element dequeued:", queue.myQueue[0])
		queue.myQueue = queue.myQueue[1:]
		return
	}
	fmt.Println("Queue is empty")

}

func GenericQueue() {

	queue := Queue[int]{
		size:    5,
		myQueue: []int{},
	}
	queue.dequeue()
	queue.enqueue(5)
	queue.enqueue(15)
	queue.enqueue(25)
	queue.enqueue(35)
	queue.dequeue()
	queue.enqueue(45)
	queue.enqueue(55)
	queue.enqueue(65)

}
