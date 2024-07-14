package main

import "fmt"

/*
	An adapter acts as a bridge between two incompatible interfaces, making them work together. This pattern involves a single class, known as the adapter,
	which is responsible for joining functionalities of independent or incompatible interfaces.

For example:
Let’s say you have two friends, one who speaks only English and another who speaks only French. You want them to communicate, but there’s a language barrier.
You act as an adapter, translating messages between them. Your role allows the English speaker to convey messages to you, and you convert those messages into French for the other person

Another example would be the ATM (adapter) which expects dollars and return rupees
*/

type Client interface {
	payInDollars(int) int
}

func main() {

	indianClient := IndianClient{}
	adapter := ExchangeAdapter{client: &indianClient}
	fmt.Println(adapter.payInDollars(1250))

}
