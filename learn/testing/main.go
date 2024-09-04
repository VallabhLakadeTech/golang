/*
Licence
*/

// Package declaration
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VallabhLakadeTech/golang/learn/testing/externalapi"
)

// Import other packages

// Global variables (exported/nonexported) - specific to the file. Common ones to be moved at utils, constants or commonhelper file

// Structs and other types - specific to the file. Common ones to be moved at utils or commonhelper file

// Functions
func main() {

	client := externalapi.APIClient{
		Client: &http.Client{},
	}
	url := "https://jsonplaceholder.typicode.com/posts"

	type POST struct {
		Title  string
		Body   string
		UserId int
		// id     int
	}
	post := POST{
		Title:  "I am title",
		Body:   "I am body",
		UserId: 150,
	}
	data, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Some error in marshalling", err)
		return
	}
	response, err := client.PostData(url, data)
	if err != nil {
		fmt.Println("Error received from PostData :: ", err)
		return
	}
	fmt.Println("Response: ", string(response))
}
