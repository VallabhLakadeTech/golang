package basic

import (
	"fmt"

	"github.com/VallabhLakadeTech/golang/grpc/my-protobuf/protogen/basic"
)

func PrintHello() {
	hello := basic.Hello{
		Name: "Vallabh",
	}

	fmt.Println("Hello world: ", hello)
}
