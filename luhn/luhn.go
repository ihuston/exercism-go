// Package luhn provides calculation of the Luhn checksum.
package luhn

import "strings"
import "unicode"

// Valid checks whether a string contains a number valid per the Luhn algorithm.
func Valid(input string) bool {
	cleaned := strings.Replace(input, " ", "", -1)
	if len(cleaned) < 2 {
		return false
	}

	sum := 0
	for n := len(cleaned) - 1; n >= 0; n-- {
		char := rune(cleaned[n])
		if !unicode.IsDigit(char) {
			return false
		}
		value := int(char - '0')
		if (len(cleaned)-n-1)%2 == 0 {
			sum += value
		} else {
			value2 := value * 2
			if value2 > 9 {
				sum += value2 - 9
			} else {
				sum += value2
			}
		}
	}
	return sum%10 == 0
}
