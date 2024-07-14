package main

type TomatoTopping struct {
	pizza IPizza
}

func (tomatoTopping *TomatoTopping) getPrice() int {
	return tomatoTopping.pizza.getPrice() + 10
}
