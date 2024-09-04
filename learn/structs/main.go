package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type Address struct {
	Street string
	City   string
	Zip    string
	a      int
	// printAddress1 func(city, street, zip string)	//Allowed but not recommended
}

func (addr Address) printAddress() {
	fmt.Println("Actual address: ", addr.City, addr.Street, addr.Zip)
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   Address  // Struct as a field
	Friends   []string // Slice as a field
}

func (person Person) printAddress() {
	fmt.Println("Persons address: ", person.Address.City, person.Address.Street, person.Address.Zip)
}

func main() {

	addr := Address{
		Street: "street",
		City:   "city",
		Zip:    "zip",
		// printAddress1: func(city, street, zip string) {
		// 	fmt.Println("Actual address: ", city, street, zip)

		// },
	}

	person := Person{
		FirstName: "vallabh",
		LastName:  "lakade",
		Age:       25,
		Address:   addr,
	}

	person.printAddress()
	addr.printAddress()

	type Person1 struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Age       int    `json:"age"`
	}

	p := Person1{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// Encode to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error in unmarshalling person struct")
		return
	}
	fmt.Println(string(jsonData))

	// Anonymous structs
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

	polymorphism()
}

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	length, width float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printArea(s shape) {
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
}

func polymorphism() {
	fmt.Println("Inside polymorphism ")
	c := circle{radius: 3.0}
	var s shape
	s = c
	fmt.Printf("At 126 :: %T\n", s)
	r := rectangle{length: 2.5, width: 1.5}
	s = r

	fmt.Printf("At 130 :: %T\n", r)
	printArea(c)
	printArea(r)
}
