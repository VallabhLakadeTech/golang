//go:build customtag
// +build customtag

package main

import "fmt"

func main() {
	// create build with customtag. go build -tags=customtag
	fmt.Println("I will run only on environments with customtag")
}
