package encoder

import (
	"fmt"
	"math"
	"strings"
)

var (
	ARABIC_NUMBERS = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	ROMAN_NUMBERS  = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

// A tuple (pair) holding our numeric value and its roman counterpart.
type NumeralTuple struct {
	arabic int
	roman  string
}

// Encodes an arabic integer to roman numeral string.
// Returns empty string and an error for invalid inputs.
func Encode(x int) (string, error) {

	if x < 0 {
		return "", fmt.Errorf("Input must be >= 0, %d given", x)
	}

	// See: https://en.wikipedia.org/wiki/Roman_numerals#Zero
	if x == 0 {
		return "nulla", nil
	}

	result := []string{}
	numerals, _ := zip(ARABIC_NUMBERS, ROMAN_NUMBERS)

	var quotient int

	for _, numeral := range numerals {
		quotient, x = divmod(x, numeral.arabic)
		result = append(result, strings.Repeat(numeral.roman, quotient))
	}

	return strings.Join(result, ""), nil
}

// Zips two lists together assuming they are the same length
func zip(a []int, b []string) ([]NumeralTuple, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("Given lists have different length")
	}

	result := make([]NumeralTuple, len(a))

	for i, e := range a {
		result[i] = NumeralTuple{e, b[i]}
	}

	return result, nil
}

// Returns (a / b, a % b) as a tuple.
func divmod(a int, b int) (int, int) {
	quotient := float64(a / b)
	remainder := math.Mod(float64(a), float64(b))

	return int(quotient), int(remainder)
}
