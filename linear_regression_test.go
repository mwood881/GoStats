package main

import (
	"math"
	"testing"
)

func TestLinearRegression(t *testing.T) {
	tests := []struct {
		name    string
		x       []float64
		y       []float64
		want    []float64
		wantErr bool
	}{
		{
			name: "Valid input",
			x:    []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			y:    []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
			want: []float64{3.00, 0.50}, // Expected coefficients (Intercept, Slope)
		},
		{
			name:    "Mismatched lengths",
			x:       []float64{10, 8, 13},
			y:       []float64{8.04, 6.95},
			wantErr: true,
		},
		{
			name: "Single point",
			x:    []float64{10},
			y:    []float64{8.04},
			want: []float64{8.04, 0}, // Expected coefficients (Intercept, Slope)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := linearRegression(tt.x, tt.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("linearRegression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !equalSlicesWithTolerance(got, tt.want, 1e-2) {
				t.Errorf("linearRegression() = %v, want %v", got, tt.want)
			}
		})
	}
}

// equalSlicesWithTolerance checks if two slices are equal within a tolerance
func equalSlicesWithTolerance(a, b []float64, tol float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > tol {
			return false
		}
	}
	return true
}

func BenchmarkLinearRegression(b *testing.B) {
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := linearRegression(x, y)
		if err != nil {
			b.Fatalf("Error during regression: %v", err)
		}
	}
}
