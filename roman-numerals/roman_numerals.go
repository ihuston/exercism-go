// Package romannumerals provides tools to transform Roman Numerals.
package romannumerals

import "errors"
import "sort"

var conversions = map[int]string{
    3000: "MMM",
    2000: "MM",
    1000: "M",
    900: "CM",
    800: "DCCC",
    700: "DCC",
    600: "DC",
    500: "D",
    400: "CD",
    300: "CCC",
    200: "CC",
    100: "C",
    90: "XC",
    80: "LXXX",
    70: "LXX",
    60: "LX",
    50: "L",
    40: "XL",
    30: "XXX",
    20: "XX",
    10: "X",
    9: "IX",
    8: "VIII",
    7: "VII",
    6: "VI",
    5: "V",
    4: "IV",
    3: "III",
    2: "II",
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
        if arabic - k >= 0 {
            result += conversions[k]
            arabic -= k
        }

        if arabic < 0 {
            break
        }
    }

    return result, nil
}

