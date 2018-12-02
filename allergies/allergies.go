// Package allergies provides tools to interpret allergy tests.
package allergies

import "sort"

var allergens = map[string]uint {
    "eggs": 0,
    "peanuts": 1,
    "shellfish": 2,
    "strawberries": 3,
    "tomatoes": 4,
    "chocolate": 5,
    "pollen": 6,
    "cats": 7,
}

// Allergies determines a list of allergens from a test score.
func Allergies(score uint) []string {
    result := []string{}
    for allergen, _ := range allergens {
        if AllergicTo(score, allergen) {
            result = append(result, allergen)
        }
    }
    sort.Strings(result)
    return result
}

// AllergicTo tests whether a score indicates an allergy to a specific allergen.
func AllergicTo(score uint, allergen string) bool {
    shift := allergens[allergen]
    if (score >> shift) % 2 == 1 {
        return true
    }
    return false
}
