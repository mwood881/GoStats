package main

import (
	"testing"
)

func TestRunRegression(t *testing.T) {
	x := []float64{8.04, 8.14, 8.74, 8.77, 9.26}
	y := []float64{6.58, 5.76, 7.71, 8.84, 8.47}

	expectedSlope := 0.5000
	expectedIntercept := 3.0000

	slope, intercept, err := RunRegression(x, y)
	if err != nil {
		t.Fatalf("Unexpected error in regression: %v", err)
	}

	if slope != expectedSlope {
		t.Errorf("Expected slope %v, but got %v", expectedSlope, slope)
	}
	if intercept != expectedIntercept {
		t.Errorf("Expected intercept %v, but got %v", expectedIntercept, intercept)
	}
}

func BenchmarkRunRegression(b *testing.B) {
	x := []float64{8.04, 8.14, 8.74, 8.77, 9.26}
	y := []float64{6.58, 5.76, 7.71, 8.84, 8.47}

	for i := 0; i < b.N; i++ {
		_, _, err := RunRegression(x, y)
		if err != nil {
			b.Fatalf("Unexpected error in regression: %v", err)
		}
	}
}
