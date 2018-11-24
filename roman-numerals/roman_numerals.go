// Package romannumerals provides tools to transform Roman Numerals.
package romannumerals

import "errors"
import "sort"

var conversions = map[int]string{
    1000: "M",
    900: "CM",
    500: "D",
    400: "CD",
    100: "C",
    90: "XC",
    50: "L",
    40: "XL",
    10: "X",
    9: "IX",
    5: "V",
    4: "IV",
    1: "I",
}

// ToRomanNumeral converts an Arabic notation number to Roman numerals.
func ToRomanNumeral(arabic int) (string, error) {
    if arabic <= 0 || arabic > 3000 {
        return "", errors.New("number outside range.")
    }

    var keys []int
    for k := range(conversions) {
        keys = append(keys, k)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(keys)))

    result := ""
    for _, k := range(keys) {
        for arabic >= k {
            result += conversions[k]
            arabic -= k
        }
        if arabic < 0 {
            break
        }
    }

    return result, nil
}

