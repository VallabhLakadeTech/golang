package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Address struct {
	street string
}

func (address Address) printMe() {
	fmt.Println("I am address")
}

type Book struct {
	name, author string
	price        float64
	ageOfAuthor  int
	address      Address
}

func main() {
	address := Address{
		street: "street",
	}
	book1 := Book{
		name:        "And then there were none",
		author:      "Agatha Criestie",
		ageOfAuthor: 100,
		address:     address,
	}

	book1.address.printMe()

	dog := struct {
		name   string
		isGood bool
	}{
		"abc",
		true,
	}
	fmt.Println(dog)

	type Person struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Age       int    `json:"age"`
	}

	p := Person{
		FirstName: "Vallabh",
		LastName:  "Lakade",
		Age:       20,
	}

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("some error in marshalling", err)
		return
	}
	fmt.Println(string(data))

	newPerson := Person{}
	err = json.Unmarshal(data, &newPerson)
	if err != nil {
		fmt.Println("some error in unmarshalling", err)
		return
	}
	fmt.Println(newPerson)

	p1 := Person{
		FirstName: "Vallabh",
		LastName:  "Lakade",
		Age:       20,
	}

	p2 := p

	if reflect.DeepEqual(p1, p2) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")

	}
	sort.
	// currencies := map[string]float64{
	// 	"usa":  83.0,
	// 	"euro": 97.0,
	// 	"abc":  0.0,
	// }

	// // var currencies map[string]float64
	// // currencies["usa"] = 12.2
	// fmt.Println(currencies["abc"])
	// if value, ok := currencies["euro"]; !ok {
	// 	fmt.Println("key doesnt exist")
	// } else {
	// 	fmt.Println("key  exist", value)
	// }

	// for key, value := range currencies {
	// 	fmt.Println("Key, value:", key, value)
	// }

	// delete(currencies, "abc")
	// fmt.Println(currencies)

	// currenciesCopy := currencies

	// currenciesCopy["euro"] = 25.0
	// fmt.Println(currencies)

	// a1 := map[string]string{"A": "a"}
	// a2 := map[string]string{"A": "a"}
	// s1 := fmt.Sprintf("%s", a1)
	// s2 := fmt.Sprintf("%s", a2)

	// if s1 == s2 {
	// 	fmt.Println("equal")
	// } else {
	// 	fmt.Println("not equal")
	// }
}
