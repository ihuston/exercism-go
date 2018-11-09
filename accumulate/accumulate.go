// Package accumulate contains tools for performing operations on a collection.
package accumulate

// Accumulate applys an operation to each element of a collection.
func Accumulate(collection []string, operation func(x string) string) []string {
	var result []string
	for _, item := range collection {
		result = append(result, operation(item))
	}
	return result
}
