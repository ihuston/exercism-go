// Package secret provides tools for decoding a secret handshake.
package secret

// Handshake decodes a secret handshake code.
func Handshake(code uint) []string {
	if code > 31 {
		return []string{}
	}
	result := make([]string, 0)

	if (code % 2) != 0 {
		result = append(result, "wink")
	}

	if (code>>1)%2 != 0 {
		result = append(result, "double blink")
	}

	if (code>>2)%2 != 0 {
		result = append(result, "close your eyes")
	}

	if (code>>3)%2 != 0 {
		result = append(result, "jump")
	}

	if (code>>4)%2 != 0 {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
	}

	return result
}
