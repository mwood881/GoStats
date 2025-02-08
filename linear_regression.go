package main

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

// linearRegression performs a simple linear regression (least squares).
func linearRegression(x []float64, y []float64) ([]float64, error) {
	if len(x) != len(y) {
		return nil, errors.New("x and y must have the same length")
	}

	// Create the design matrix with an intercept term
	X := mat.NewDense(len(x), 2, nil)
	for i := 0; i < len(x); i++ {
		X.Set(i, 0, 1)    // Intercept term
		X.Set(i, 1, x[i]) // Independent variable
	}

	// Create the response vector
	Y := mat.NewVecDense(len(y), y)

	// Solve for coefficients using least squares
	var coeffs mat.VecDense
	err := coeffs.SolveVec(X, Y)
	if err != nil {
		return nil, err
	}

	// Extract coefficients
	intercept := coeffs.AtVec(0)
	slope := coeffs.AtVec(1)

	return []float64{intercept, slope}, nil
}
