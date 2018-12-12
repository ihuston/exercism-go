// Package wordcount provides tools to count words.
package wordcount

import "unicode"
import "strings"

// Frequency holds the frequency of words in a phrase.
type Frequency map[string]int

// WordCount returns the count of words in a provided phrase.
func WordCount(phrase string) Frequency {
	words := splitter(phrase)
	counts := map[string]int{}
	for _, w := range words {
		counts[w]++
	}
	return counts
}

func splitter(phrase string) []string {
	words := []string{}
	var w strings.Builder
	for _, r := range phrase {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'' {
			w.WriteRune(unicode.ToLower(r))
		} else {
			if w.Len() != 0 {
				words = append(words, strings.Trim(w.String(), "'\""))
			}
			w.Reset()
		}
	}
	if w.Len() != 0 {
		words = append(words, strings.Trim(w.String(), "'\""))
	}
	return words
}
