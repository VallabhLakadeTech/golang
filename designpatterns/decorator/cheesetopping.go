package main

type CheeseTopping struct {
	pizza IPizza
}

func (cheeseTopping *CheeseTopping) getPrice() int {
	return cheeseTopping.pizza.getPrice() + 15
}
