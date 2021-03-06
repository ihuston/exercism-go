// Package leap provides leap year calculation tools.
package leap

// IsLeapYear determines whether the provided year is a leap year.
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	return true
}

