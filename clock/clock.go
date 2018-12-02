// Package clock provides a simple clock.
package clock

import "fmt"

// Clock represents a 24 hour clock.
type Clock int

func clockMod(input int) Clock {
	for input < 0 {
		input += 24 * 60
	}
	return Clock(input % (24 * 60))
}

// New returns a new clock.
func New(hour, minute int) Clock {
	return clockMod(hour*60 + minute)
}

// String helper function.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add minutes to a clock.
func (c Clock) Add(minutes int) Clock {
	return clockMod(int(c) + minutes)
}

// Subtract minutes from a clock.
func (c Clock) Subtract(minutes int) Clock {
	return clockMod(int(c) - minutes)
}
