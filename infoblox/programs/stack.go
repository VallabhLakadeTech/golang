package programs

import "fmt"

//Implement Stack DS using Generics / Interface

type AllowdTypes interface {
	int | float64 | string
}

type Stack[T AllowdTypes] struct {
	myStack []T
	size    int
}

func (stack *Stack[T]) pop() {
	currentSize := len(stack.myStack)
	if currentSize != 0 {
		fmt.Println("Last Element: ", stack.myStack[currentSize-1])
		stack.myStack = stack.myStack[:currentSize-1]
		return
	}
	fmt.Println("Stack is empty")
}

func (stack *Stack[T]) push(element T) {
	if len(stack.myStack) != stack.size {
		stack.myStack = append(stack.myStack, element)
		fmt.Println("New stack: ", stack.myStack)
		return
	}
	fmt.Println("Stack is full")
}

func GenericStack() {

	// stack := Stack[int]{
	// 	size:    5,
	// 	myStack: []int{},
	// }
	// stack.pop()
	// stack.push(5)
	// stack.push(10)
	// stack.push(15)
	// stack.push(20)
	// stack.pop()
	// stack.push(25)
	// stack.push(30)
	// stack.push(35)

	stack := Stack[string]{
		size:    5,
		myStack: []string{},
	}
	stack.pop()
	stack.push("5-1")
	stack.push("10-1")
	stack.push("15-1")
	stack.push("20-1")
	stack.pop()
	stack.push("25-1")
	stack.push("30-1")
	stack.push("35-1")
	stack.push("35-1")
	stack.pop()

}
