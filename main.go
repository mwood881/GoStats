package main

import (
	"fmt"
	"log"
	"time"
)

// Only have one main function in folder
// Do not include func main in linear regression
func main() {
	// Define the Anscombe data
	// Include x values like from R and Python files
	//keep type same as y
	x := [][]float64{
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
	}

	// Include y values like from R and Python files
	// y needs float sinvce decimal numbers
	y := [][]float64{
		{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
		{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
		{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
		{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
	}

	// mark the time it takes to run
	start := time.Now()

	// Fit linear regression models for each set of data
	for i := 0; i < 4; i++ {
		// Get the data for the current set
		xData := x[i]
		yData := y[i]

		// Perform linear regression (calling the function defined in linear_regression.go)
		// create file linear regression but do not run individually
		// run this: go run main.go linear_regression.go instead then % go test -bench
		coeffs, err := linearRegression(xData, yData)
		if err != nil {
			log.Fatalf("Error during regression for Set %d: %v", i+1, err)
		}

		// Output Intercept and Slope for the regression coefficients
		fmt.Printf("Regression model coefficients for Set %d: Intercept: %.2f, Slope: %.2f\n", i+1, coeffs[0], coeffs[1])
	}

	// Output the total time and memory
	fmt.Printf("Total time: %v\n", time.Since(start))
}
