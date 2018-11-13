// Package isogram provides tools for checking whether a string contains repeated letters.
package isogram

import "strings"

// IsIsogram checks whether a string contains any repeated letters.
func IsIsogram(input string) bool {
	cleaned := strings.ToLower(input)
	cleaned = strings.Replace(cleaned, " ", "", -1)
	cleaned = strings.Replace(cleaned, "-", "", -1)
	seen := make(map[rune]bool)
	for _, r := range cleaned {
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
