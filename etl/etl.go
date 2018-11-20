// Package etl provides tools for transforming data structures
package etl

import "strings"

// Transform from a list of letters per score into a single letter per score table.
func Transform(scoresToLetters map[int][]string) map[string]int {
	result := make(map[string]int)
	for score, letters := range scoresToLetters {
		for _, s := range letters {
			result[strings.ToLower(s)] = score
		}
	}

	return result
}
