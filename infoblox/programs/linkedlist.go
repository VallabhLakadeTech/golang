package programs

import (
	"fmt"
)

type AllowdDataTypes interface {
	int | float64
}

type Node[T AllowdDataTypes] struct {
	data T
	next *Node[T]
}

type LinkedList[T AllowdDataTypes] struct {
	head *Node[T]
}

func (ll *LinkedList[T]) add(data T) {
	newLL := Node[T]{
		data: data,
	}
	if ll.head != nil {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = &newLL
		return
	}
	ll.head = &newLL
}

func (ll *LinkedList[T]) remove(data T) {

	if ll.head == nil {
		fmt.Println("Empty linkedlist")
		return
	}
	current := ll.head
	previous := ll.head
	for current.next != nil {
		if current.data == data {
			if current == ll.head {
				ll.head = current.next
			} else {
				previous.next = current.next
			}
			return
		}
		previous = current
		current = current.next
	}

}

func (ll *LinkedList[T]) print() {
	current := ll.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func GenericLinkedList() {

	ll := LinkedList[int]{}
	ll.add(1)
	ll.add(2)
	ll.add(3)
	ll.print()
	ll.remove(2)
	ll.print()
	ll.add(4)
	ll.add(5)
	ll.print()
}
