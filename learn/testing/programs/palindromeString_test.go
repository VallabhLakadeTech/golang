package programs

import (
	"fmt"

	"github.com/stretchr/testify/assert"

	"testing"
)

// go get github.com/go-delve/delve/cmd/dlv
func TestReverse(t *testing.T) {
	isPalindrome := Reverse("lil")
	assert.True(t, isPalindrome)
}

// go test -v : runs all test cases
// go test -run TestFileName -v : runs testcases from a file
// go test -run TestFileName/Test_Case -v : runs a single test case from a test suite of a test file
func TestTwoPointer(t *testing.T) {

	type TestSuite struct {
		expectedValue bool
		name          string
		inputStr      string
	}

	suits := []TestSuite{
		{
			inputStr:      "lilli",
			name:          "first test case",
			expectedValue: false,
		},
		{
			inputStr:      "",
			name:          "second test case",
			expectedValue: true,
		},
		{
			inputStr:      "lil2",
			name:          "third test case",
			expectedValue: false,
		},
	}

	for i := 0; i < len(suits); i++ {
		t.Run(suits[i].name, func(t *testing.T) {
			returnedValue := Reverse(suits[i].inputStr)
			assert.Equal(t, returnedValue, suits[i].expectedValue)
		})
	}

}

// go test -bench=.
// go test -bench BenchmarkReverse
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("lil")
	}
}

/*
Output:
go test -bench=.
(Name-GOMAXPROCS No. of parallel processes)		Number of iterations	Avg time per operation
BenchmarkReverse-12             				13868258                83.16 ns/op

go test -bench=. -benchtime 5s -benchmem

(Name-GOMAXPROCS No. of parallel processes)		Number of iterations	Avg time per operation	No. of bytes per operation		No. of allocation per operation
BenchmarkReverse-12             				71425720                82.72 ns/op            	6 B/op          				3 allocs/op

go test -bench=. -run=xxxx : Skip running any test cases

go test -bench BenchmarkReverse -benchmem -memprofile mem1.pprof -cpuprofile  cpu1.pprof

Install brew install graphviz for web UI
go tool pprof mem1.pprof
web
top 20
list NAME
*/
func BenchmarkTwoPointers(b *testing.B) {

	strList := []string{"madam", "racecar", "pointer", "refer", "after", "noon", "stats", "civic", "mathematics"}
	for i := 0; i < len(strList); i++ {
		b.Run(fmt.Sprintf("TwoPointers-%d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				twoPointer(strList[i])
			}
		})
	}

}
