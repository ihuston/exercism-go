// Package proverb provides tools for creating proverbs.
package proverb

// Proverb constructs a proverb from the provided strings.
func Proverb(rhyme []string) []string {
	response := []string{}
	for i := 1; i < len(rhyme); i++ {
		response = append(response, "For want of a "+rhyme[i-1]+" the "+rhyme[i]+" was lost.")
	}

	if len(rhyme) > 0 {
		response = append(response, "And all for the want of a "+rhyme[0]+".")
	}
	return response
}
