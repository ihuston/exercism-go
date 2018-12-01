// Package perfect provides classification of aliquot sums.
package perfect

import "errors"

// Classification is perfect, abundant or deficient.
type Classification int

// These constants are the Classifications.
const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

// ErrOnlyPositive when not a positive integer.
var ErrOnlyPositive = errors.New("not a positive integer")

// Classify as a perfect, abundant, or deficient aliquot sum.
func Classify(input int64) (Classification, error) {
	if input <= 0 {
		return Classification(-1), ErrOnlyPositive
	}
	sum := int64(0)
	for i := int64(1); i < input; i++ {
		if input%i == 0 {
			sum += i
		}
	}
	if sum > input {
		return ClassificationAbundant, nil
	} else if sum == input {
		return ClassificationPerfect, nil
	}
	return ClassificationDeficient, nil
}
