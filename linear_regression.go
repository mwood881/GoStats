package main

// when importing do not use github in the beginning of the link
// when you do this it somehow always produces an error
import (
	"gonum.org/v1/gonum/mat"
)

// linearRegression performs a simple linear regression (least squares).
func linearRegression(x []float64, y []float64) ([]float64, error) {
	// Add constant  to the independent variable
	// this is the intercept variable
	X := make([][]float64, len(x))
	for i := 0; i < len(x); i++ {
		X[i] = []float64{1, x[i]} // 1 for intercept
	}

	// matrix needed to be created
	XMat := mat.NewDense(len(X), 2, nil)
	for i, row := range X {
		XMat.SetRow(i, row)
	}
	YMat := mat.NewVecDense(len(y), y)

	// Solve for coefficients using least squares
	// find coefficients
	var coeffs mat.VecDense
	err := coeffs.SolveVec(XMat, YMat)
	if err != nil {
		return nil, err
	}
	// Return intercept and slope
	return []float64{coeffs.At(0, 0), coeffs.At(1, 0)}, nil
}
