package main

import "testing"

var area float64

// go test -benchmem -run=^$ -bench ^BenchmarkAreaOfCircle$ github.com/VallabhLakadeTech/golang/testcases
func BenchmarkAreaOfCircle(b *testing.B) {

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {

		area = areaOfCircle(10)

	}

}

func BenchmarkAreaOfCircleAllocs(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		areaOfCircle(10)
	}
}

func BenchmarkAreaOfCircleParallel(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		if p.Next() {
			areaOfCircle(10)
		}
	})
}
