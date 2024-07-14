//go:build darwin && !windows && !linux && !customtag
// +build darwin,!windows,!linux,!customtag

package main

import "fmt"

func main() {
	fmt.Println("I will run only on Mac machines", config_value)
}
