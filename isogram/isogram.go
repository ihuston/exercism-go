// Package isogram provides tools for checking whether a string contains repeated letters.
package isogram

import "unicode"

// IsIsogram checks whether a string contains any repeated letters.
func IsIsogram(input string) bool {
	seen := make(map[rune]bool)
	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		lowered := unicode.ToLower(r)
		if seen[lowered] {
			return false
		}
		seen[lowered] = true
	}
	return true
}
