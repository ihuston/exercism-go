// Package cryptosquare provides encoding tools.
package cryptosquare

import (
    "unicode"
    "math"
)

// Encode plain text using a square code.
func Encode(plain string) string {
    if len(plain) == 0 {
        return ""
    }
    cleaned := clean(plain)
    r, c := shape(len(cleaned))
    split := blocks(cleaned, r, c)
    result := ""
    for i:=0; i < c; i++ {
        for j:=0; j < r; j++ {
            result += string(split[j][i])
        }
        result += " "
    }
    return result[:(r+1)*c-1]
}

func clean(dirty string) string {
    cleaned := ""
    for _, r := range(dirty) {
        if !(unicode.IsLetter(r) || unicode.IsDigit(r)) {
            continue
        }
        cleaned += string(unicode.ToLower(r))
    }
    return cleaned
}

func shape(length int) (rows int, cols int) {
    root := math.Sqrt(float64(length))
    cols = int(math.Ceil(root))
    rows = int(math.Round(root))
    return
}

func blocks(input string, r int, c int) []string {
    split := []string{}
    for len(input) < r*c {
        input += " "
    }
    for i:=0; i < r; i++ {
        split = append(split, input[i*c:(i+1)*c])
    }
    return split
}
