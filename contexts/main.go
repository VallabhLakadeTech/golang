package main

import (
	"context"
	"fmt"
)

type myPrivateKey string

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	withCancel(ctx)

	ctx1 := context.WithValue(context.Background(), myPrivateKey("some_key"), "some_value")
	withValue(ctx1)

	// ctx, cancel := context.WithDeadline(time.Second*10)

	// context.WithTimeout()
}

func withCancel(ctx context.Context) {
	fmt.Println("In withContext")
}

func withValue(ctx context.Context) {
	fmt.Println("In withValue", ctx.Value(myPrivateKey("some_key")))
}

func useContext(ctx context.Context, ch chan int) error {

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-ch:
		fmt.Println("Something")
	}

}
