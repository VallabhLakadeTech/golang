package main

import "fmt"

func main() {

	fmt.Println(areaOfCircle(3.5))

}

func areaOfCircle(radius float64) float64 {
	return 3.14 * radius * radius
}
