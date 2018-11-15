// Package diffsquares provides functions for the difference of squares of natural numbers.
package diffsquares

import "math"

// SquareOfSum returns the square of the sum of natural numbers up to the specified integer.
func SquareOfSum(n int) int {
	return int(math.Pow(float64(n*(n+1.0)/2), 2))
}

// SumOfSquares returns the sum of the squares of natural numbers up to the specified integer.
func SumOfSquares(n int) int {
	return int((n * (n + 1) * (2*n + 1)) / 6)
}

// Difference returns the difference of the sum of squares and square of the sums of natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
