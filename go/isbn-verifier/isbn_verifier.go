package isbnverifier

import (
	"regexp"
	"strings"
)

func IsValidISBN(isbn string) bool {
	clearedIsbn := strings.ReplaceAll(isbn, "-", "")

	isbnRegex := regexp.MustCompile(`^\d{9}[\dX]$`)

	if !isbnRegex.MatchString(clearedIsbn) {
		return false
	}

	return calculateIfValid(clearedIsbn)
}

func calculateIfValid(isbn string) bool {
	// (d₁ * 10 + d₂ * 9 + d₃ * 8 + d₄ * 7 + d₅ * 6 + d₆ * 5 + d₇ * 4 + d₈ * 3 + d₉ * 2 + d₁₀ * 1) mod 11 == 0

	sum := 0

	for i := 0; i < len(isbn); i++ {
		multiplyer := len(isbn) - i

		var digit int
		if isbn[i] == 'X' {
			digit = 10
		} else {
			digit = int(isbn[i] - '0')
		}

		sum += digit * multiplyer
	}

	// digit := int(id[i] - '0')
	return sum%11 == 0
}
