package strand

import "strings"

var coding = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA transcribes from DNA sequences to RNA ones.
func ToRNA(dna string) string {
	if dna == "" {
		return ""
	}
	var s strings.Builder
	for _, r := range dna {
		s.WriteRune(coding[r])
	}
	return s.String()
}
