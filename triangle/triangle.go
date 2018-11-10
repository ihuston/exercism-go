// Package triangle provides a test of triangle types.
package triangle

import "math"

// Kind specifies a type of triangle.
type Kind int

// Constants for different types of triangles.
const (
	NaT Kind = iota // not a triangle
	Equ // equilateral
	Iso // isosceles
	Sca // scalene
)

// KindFromSides returns the Kind of triangle based on length of the sides.
func KindFromSides(a, b, c float64) Kind {
	if !((a > 0) && (b > 0) && (c > 0)) {
		return NaT
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}
	if (a+b < c) || (b+c < a) || (c+a < b) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if (a == b) || (b == c) || (c == a) {
		return Iso
	}
	return Sca
}
