// Package raindrops provides responses based on factorisation.
package raindrops

import "strconv"

// Convert takes a number and responds based on whether 3, 5, 7 are factors.
func Convert(number int) string {
	response := ""
	if number%3 == 0 {
		response += "Pling"
	}
	if number%5 == 0 {
		response += "Plang"
	}
	if number%7 == 0 {
		response += "Plong"
	}
	if response == "" {
		response = strconv.Itoa(number)
	}
	return response
}
