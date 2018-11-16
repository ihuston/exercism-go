// Package luhn provides calculation of the Luhn checksum.
package luhn

import "strconv"
import "strings"

// Valid checks whether a string contains a number valid per the Luhn algorithm.
func Valid(input string) bool {
	cleaned := strings.Replace(input, " ", "", -1)
	if len(cleaned) < 2 {
		return false
	}

	_, err := strconv.ParseInt(cleaned, 10, 64)
	if err != nil {
		return false
	}

	evens := 0
	odds := 0
	for n := range cleaned {
		index := len(cleaned) - n - 1
		digit := cleaned[index]

		value := int(digit) - int('0')
		if n%2 == 0 {
			evens += value
		} else {
			value2 := value * 2
			if value2 > 9 {
				odds += value2 - 9
			} else {
				odds += value2
			}
		}
	}

	if (odds+evens)%10 == 0 {
		return true
	}
	return false
}
