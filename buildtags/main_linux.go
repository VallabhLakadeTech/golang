//go:build linux
// +build linux

package main

import "fmt"

func main() {
	fmt.Println("I will run only on linux environment", config_value)
}
