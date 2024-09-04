package main

import (
	"fmt"
)

// type parameter
// type constraint

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(add(3.5, 4.3))

	type num int
	a := num(10)
	b := num(10)
	fmt.Println(add(a, b))

	// printMe(1)
	// printMe("Vallabh")

	emp := Employee[float32]{
		name:   "Vallabh",
		age:    26,
		salary: 20,
	}
	emp.getSal()
}

type AllowedTypes interface {
	~int | ~float64
}

func add[T AllowedTypes](a, b T) T {
	return a + b
}

type Employee[T int | float32] struct {
	name   string
	age    int
	salary T
}

func (emp Employee[T]) getSal() T {
	return emp.salary
}
