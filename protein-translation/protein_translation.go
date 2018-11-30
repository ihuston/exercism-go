// Package protein provides tools for protein translation.
package protein

import "errors"

var STOP error = errors.New("stop code")
var ErrInvalidBase error = errors.New("invalid base")

var codons = map[string]string{
    "AUG": "Methionine",
    "UUU": "Phenylalanine",
    "UUC": "Phenylalanine",
    "UUA": "Leucine",
    "UUG": "Leucine",
    "UCU": "Serine",
    "UCC": "Serine",
    "UCA": "Serine",
    "UCG": "Serine",
    "UAU": "Tyrosine",
    "UAC": "Tyrosine",
    "UGU": "Cysteine",
    "UGC": "Cysteine",
    "UGG": "Tryptophan",
}

// FromCodon translates from codons to proteins.
func FromCodon(input string) (string, error) {
    if input == "UAA" || input == "UAG" || input == "UGA" {
        return "", STOP
    }
    if _, ok := codons[input]; ok != true {
        return "", ErrInvalidBase
    }
    return codons[input], nil
}

// FromRNA translates from RNA to proteins.
func FromRNA(input string) ([]string, error) {
    results := []string{}
    for i := 0; i < len(input)/3; i++ {
        protein, err := FromCodon(input[3*i:3*(i+1)])
        if err == STOP {
            return results, nil
        }
        if err == ErrInvalidBase {
            return results, err
        }
        results = append(results, protein)
    }
    return results, nil
}
