// Package grains provides tools to calculate grains on a chessboard.
package grains

import "errors"

//import "math"

// Square calculates the number of grains on a specified square of a chessboard.
func Square(input int) (uint64, error) {
	if input <= 0 || input >= 65 {
		return 0, errors.New("calculation failed")
	}

	// More readable but slower library call (316 ns/op)
	//result := math.Pow(2.0, float64(input-1))

	// Less readable but faster bit shift approach (17.8 ns/op)
	result := 1 << uint64(input-1)
	return uint64(result), nil
}

// Total calculates the total number of grains on a whole chessboard.
func Total() uint64 {
	sum := uint64(0)
	for i := 1; i < 65; i++ {
		s, _ := Square(i)
		sum += s
	}
	return sum
}
