package main

import (
	"testing"
)

// This took forever to understand ugh
// Include a benchmark or a test to check if the x and y given produce a linear regression
func BenchmarkLinearRegression(b *testing.B) {
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	//show if there is an error or not when finding a benchmark
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// Perform linear regression for benchmarking
		_, err := linearRegression(x, y)
		if err != nil {
			b.Fatalf("Error during regression: %v", err)
		}
	}
}
