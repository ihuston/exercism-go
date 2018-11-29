// Package pythagorean
package pythagorean

import "math"

type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the range min to max inclusive.
func Range(min, max int) []Triplet {
    results := []Triplet{}
    for a := min; a <= max; a++ {
        for b := a; b <= max; b++ {
            c := math.Sqrt(float64(a*a + b*b))
            if c == math.Floor(c) && int(c) <= max {
                results = append(results, Triplet{a, b, int(c)})
            }
        }
    }
    return results
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c is equal to p.
func Sum(p int) []Triplet {
    triplets := Range(1, 1001)
    results := []Triplet{}
    for _, t := range triplets {
        if t[0] + t[1] + t[2] == p {
            results = append(results, t)
        }
    }
    return results
}
