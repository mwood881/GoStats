package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

// Dataset struct holds x and y values
type Dataset struct {
	X []float64
	Y []float64
}

// RunRegression performs linear regression and returns slope and intercept
func RunRegression(x, y []float64) (float64, float64, error) {
	// Perform linear regression using the stats package
	// LinearRegression takes two slices of float64 (x and y), not stats.Series
	regression, err := stats.LinearRegression(x, y, false)
	if err != nil {
		return 0, 0, err
	}

	// The result is of type stats.Regression, and you can access Slope and Intercept
	return regression.Slope, regression.Intercept, nil
}

func main() {
	// Example of Dataset 1 (Anscombe I)
	anscombe1 := Dataset{
		X: []float64{8.04, 8.14, 8.74, 8.77, 9.26},
		Y: []float64{6.58, 5.76, 7.71, 8.84, 8.47},
	}

	// Run regression
	slope, intercept, err := RunRegression(anscombe1.X, anscombe1.Y)
	if err != nil {
		fmt.Println("Error in regression:", err)
		return
	}

	// Output the result
	fmt.Printf("Slope: %.4f, Intercept: %.4f\n", slope, intercept)
}
