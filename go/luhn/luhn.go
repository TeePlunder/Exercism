package luhn

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	if len(id) <= 1 {
		return false
	}

	onlyDigitsRegex := regexp.MustCompile(`^[0-9]+$`)

	if !onlyDigitsRegex.MatchString(id) {
		return false
	}

	trimedId := strings.TrimSpace(id)

	doubledId := doubleEverySecond(trimedId)

	return true
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
