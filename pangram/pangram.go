// Package pangram provides tools to determine if a sentence is a pangram.
package pangram

import "unicode"

// IsPangram checks whether a sentence contains all letters of the alphabet.
func IsPangram(s string) bool {
    for _, r := range("abcdefghijklmnopqrstuvwxyz") {
        if !contains(s, r) {
            return false
        }
    }
    return true
}

func contains(input string, r rune) bool {
    for _, s := range(input) {
        if unicode.ToLower(s) == unicode.ToLower(r) {
            return true
        }
    }
    return false
}
