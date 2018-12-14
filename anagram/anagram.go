// Package anagram provides tools to detect anagrams.
package anagram

import "sort"
import "fmt"
import "strings"

// Detect checks possible anagrams of a phrase.
func Detect(subject string, candidates []string) []string {

	results := candidates[:0]
	for _, c := range candidates {
		if anagrams(c, subject) {
			results = append(results, c)
		}
	}
	return results
}

func anagrams(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	if a == b {
		return false
	}
	a = sortString(a)
	b = sortString(b)
	if a == b {
		return true
		fmt.Println(a + "==" + b)
	}
	return false
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
