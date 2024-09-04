package main

import (
	"fmt"
	"time"

	"github.com/VallabhLakadeTech/golang/learn/concurrency/topics"
)

func main() {

	// go printMe("Dog")
	//go printMe("Cat")

	// time.Sleep(1 * time.Second)

	// Waitgroup
	// topics.WaitGroup()

	//Channels
	// topics.Channels()

	// Select statement
	// topics.Select()

	// topics.AvoidingRace()

	topics.Workerpool()

}

func printMe(animal string) {

	for i := 0; i < 10; i++ {
		fmt.Println(animal)
		time.Sleep(time.Millisecond * 500)
	}
}
