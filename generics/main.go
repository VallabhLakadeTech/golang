package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func add[T ~int | float64](a, b T) T {
	return a + b
}

type Num interface {
	int | int8 | int32 | int64 | float32 | float64
}

func add1[T Num](a, b T) T {
	return a + b
}

func add2[T constraints.Ordered](a, b T) T {
	return a + b
}

// Generics on struct
// Following is called type set feature
type CustomSalary interface {
	int | float64
}

// Variables should be exported
type Employee[T CustomSalary] struct {
	Name    string
	Age     int
	Package T
}

func main() {

	fmt.Println(add(3, 4))
	fmt.Println(add(3.5, 4.25))

	fmt.Println(add1(4, 5))
	fmt.Println(add1(4, 4.25))
	fmt.Println(add1(3.5, 4.25))

	fmt.Println(add2(3.75, 4.25))

	type MyNum int
	a := MyNum(2)
	b := MyNum(3)
	fmt.Println(add(a, b))

	// Create a Map function using generics
	mappedIntValues := MapFunc([]int{2, 4, 6}, func(value int) int {
		return value * 5
	})
	fmt.Println("Mapping int values: ", mappedIntValues)

	mappedFloatValues := MapFunc([]float64{2.5, 4.5, 6.5}, func(value float64) float64 {
		return value * 5
	})
	fmt.Println("Mapping float values: ", mappedFloatValues)

	employee := Employee[float64]{
		Name:    "Vallabh",
		Age:     25,
		Package: 10000.500,
	}
	fmt.Println("Employee: ", employee)

	// Generics on Map
	type GenericMap[T comparable, V int | string] map[T]V
	student := make(GenericMap[string, int])
	student["Age"] = 25
	fmt.Println("Generic Map: ", student)
}

func MapFunc[T int | float64](values []T, mapping func(T) T) (mappedValueArr []T) {
	for _, v := range values {
		mappedValue := mapping(v)
		mappedValueArr = append(mappedValueArr, mappedValue)
	}
	return mappedValueArr
}
