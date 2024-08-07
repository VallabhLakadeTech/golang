package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v : runs all test cases
// go test -run TestFileName -v : runs testcases from a file
// go test -run TestFileName/Test_Case -v : runs a single test case from a test suite of a test file
func TestAreaOfCircle(t *testing.T) {

	// Table driven tests
	type TestSuite struct {
		Name     string
		Radius   float64
		Expected float64
	}

	testSuite := []TestSuite{
		{
			Name:     "Radius 1",
			Radius:   1,
			Expected: 3.14,
		},
		{
			Name:     "Radius 5",
			Radius:   5,
			Expected: 78.5,
		},
		{
			Name:     "Radius 10",
			Radius:   10,
			Expected: 314,
		},
		// {
		// 	Name:     "Radius 15",
		// 	Radius:   15,
		// 	Expected: 524,
		// },
	}

	for _, test := range testSuite {
		t.Run(test.Name, func(t *testing.T) {
			actual := areaOfCircle(test.Radius)
			assert.Equal(t, test.Expected, actual)
		})
	}

}
