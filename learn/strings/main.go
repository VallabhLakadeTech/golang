package main

import (
	"fmt"
	"time"
)

func main() {

	//s1 := "I am a string."
	s1 := "Héllô GÕ!"
	// rune : int32
	//string: uint8
	fmt.Println(s1)

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c ", s1[i])
	}
	s1rune := []rune(s1)
	for key, value := range s1rune {
		fmt.Printf("%v,%c\n", key, value)
	}

	variadic()

	a := func(msg string) {
		fmt.Println("Hello", msg)
	}
	a("val")

	day := 24 * time.Hour
	seconds := day.Seconds()
	fmt.Println(seconds)
}

type Details struct {
	name string
}

func (detail Details) printName() {
	fmt.Println("Name: ", detail.name)
}

func returnsomething(a, b int, c string) (int, error) {
	return 1, nil
}

func variadic() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("Inside variadic func")
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))

}

func sum(numbers ...int) (sum int) {

	for _, number := range numbers {
		sum = sum + number
	}
	return
}
