// Package lsproduct provides calculations of the largest product in a series.
package lsproduct

import "errors"
import "unicode"

// LargestSeriesProduct finds the largest product of a provided number of consequtive digits.
func LargestSeriesProduct(digitString string, span int) (int64, error) {
    if len(digitString) < span {
        return 0, errors.New("span too long")
    }
    if span < 0 {
        return 0, errors.New("span is negative")
    }

    digits, err := digitsFromString(digitString)
    if err != nil {
        return 0, errors.New("error with string")
    }
    max := int64(0)
    for start := 0; start < len(digits) - span + 1; start++ {
        product := int64(1)
        for i := 0; i < span; i++ {
            product *= digits[start+i]
        }
        if product > max {
            max = product
        }
    }
    return max, nil
}

func digitsFromString(input string) ([]int64, error) {
    result := []int64{}
    for _, r := range(input) {
        if !unicode.IsDigit(r) {
            return []int64{}, errors.New("not a digit")
        }
        result = append(result, int64(r - '0'))
    }
    return result, nil
}
