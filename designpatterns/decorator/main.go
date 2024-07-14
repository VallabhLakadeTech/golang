package main

import "fmt"

type IPizza interface {
	getPrice() int
}

// The Decorator Pattern is a structural design pattern that allows behavior to be added to individual objects, dynamically and transparently, without affecting
// the behavior of other objects from the same class. This pattern provides an alternative to subclassing for extending functionality.

func main() {

	veggiePizza := Veggie{}
	fmt.Println("Price of a Veggie pizza:", veggiePizza.getPrice())

	tomatoToppingPizza := TomatoTopping{pizza: &veggiePizza}
	fmt.Println("Price of a tomato topping pizza:", tomatoToppingPizza.getPrice())

	cheeseWithTomatoTopping := CheeseTopping{pizza: &tomatoToppingPizza}
	fmt.Println("Price of a cheese with tomato topping pizza:", cheeseWithTomatoTopping.getPrice())

}
