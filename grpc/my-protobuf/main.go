package main

import (
	"github.com/VallabhLakadeTech/golang/grpc/my-protobuf/basic"
)

func main() {

	basic.PrintHello()
	basic.PrintUser()
	basic.ProtoToUserJson()

}
