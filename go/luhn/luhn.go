package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	trimedId := strings.ReplaceAll(id, " ", "")

	if len(trimedId) <= 1 {
		return false
	}

	onlyDigitsRegex := regexp.MustCompile(`^[0-9]+$`)

	if !onlyDigitsRegex.MatchString(trimedId) {
		return false
	}

	doubledId := doubleEverySecond(trimedId)

	sumIsValid := sumIsValid(doubledId)

	return sumIsValid
}

func sumIsValid(id string) bool {
	sum := 0
	for _, element := range id {
		digit := int(element - '0')
		sum += digit
	}

	return sum%10 == 0
}

func doubleEverySecond(id string) string {
	doubledId := ""
	isPenultimatePosition := false

	for i := len(id) - 1; i >= 0; i-- {
		// trimedId[i] is a byte holding the ASCII code, not the numeric value.
		// '0'..'9' map to 48..57, so subtract '0' (48) to get the actual digit.
		digit := int(id[i] - '0')

		if !isPenultimatePosition {
			isPenultimatePosition = !isPenultimatePosition
			doubledId = strconv.Itoa(digit) + doubledId
			continue
		}

		digit *= 2

		if digit > 9 {
			digit -= 9
		}

		isPenultimatePosition = !isPenultimatePosition
		doubledId = strconv.Itoa(digit) + doubledId
	}

	return doubledId
}
