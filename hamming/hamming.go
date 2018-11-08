// Package hamming provides comparison tools.
package hamming

import "errors"

// Distance calculates the Hamming distance between two strings.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("string lengths do not match")
	}
	count := 0
	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
