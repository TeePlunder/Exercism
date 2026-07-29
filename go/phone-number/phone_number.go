package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

func Number(phoneNumber string) (string, error) {
	nonDigitRegex := regexp.MustCompile(`\D`)

	cleanedPhoneNumber := nonDigitRegex.ReplaceAllString(phoneNumber, "")

	cleanedPhoneNumberLength := len(cleanedPhoneNumber)

	if cleanedPhoneNumberLength == 11 && cleanedPhoneNumber[0] == '1' {
		cleanedPhoneNumber = cleanedPhoneNumber[1:]
	}

	// checking for N in NXX NXX-XXXX (N is in between 2 and 9)
	firstN := toDigit(cleanedPhoneNumber[0])
	secondN := toDigit(cleanedPhoneNumber[3])

	if !isValidN(firstN) || !isValidN(secondN) {
		return "", errors.New("invalid N")
	}

	return cleanedPhoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	cleanedPhoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return cleanedPhoneNumber[:3], nil
}

func Format(phoneNumber string) (string, error) {
	cleanedPhoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	formatted := fmt.Sprintf("(%s) %s-%s", cleanedPhoneNumber[0:3], cleanedPhoneNumber[3:6], cleanedPhoneNumber[6:])

	return formatted, nil
}

func toDigit(element byte) int {
	return int(element - '0')
}

func isValidN(n int) bool {
	return 2 <= n || n <= 9
}
