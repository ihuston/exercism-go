// Package atbash provides tools for substitution ciphers.
package atbash

import "strings"
import "unicode"

var code = strings.NewReplacer(
	"a", "z",
	"b", "y",
	"c", "x",
	"d", "w",
	"e", "v",
	"f", "u",
	"g", "t",
	"h", "s",
	"i", "r",
	"j", "q",
	"k", "p",
	"l", "o",
	"m", "n",
	"n", "m",
	"o", "l",
	"p", "k",
	"q", "j",
	"r", "i",
	"s", "h",
	"t", "g",
	"u", "f",
	"v", "e",
	"w", "d",
	"x", "c",
	"y", "b",
	"z", "a",
)

// Atbash transforms a string using a simple cipher.
func Atbash(input string) string {
	// prep
	var p strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			p.WriteRune(r)
		}
	}
	//decode
	full := code.Replace(strings.ToLower(p.String()))
	//chunker
	var s strings.Builder
	for i, r := range full {
		if (i > 0) && (i%5 == 0) {
			s.WriteRune(' ')
		}
		s.WriteRune(r)
	}
	return s.String()
}
