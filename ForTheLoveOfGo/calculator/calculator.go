// Package calculator does simple calculations.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding
// them together
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers a and b, and
// returns the result of multiplying
func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of negative number")
	}
	return math.Sqrt(a), nil
}
