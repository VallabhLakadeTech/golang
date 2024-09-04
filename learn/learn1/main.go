package main

import (
	"fmt"
)

// var Age = 10 //exported

func main() {

	// var age = 10
	fmt.Println(basic.Age)

	var a int
	var b float64
	var c string // ""
	var d bool
	fmt.Println(a, b, c, d)

}

func myfunc() (int, int) {
	return 1, 2
}
