package raindrops

import "strconv"

func Convert(number int) string {
	addPling := number%3 == 0
	addPlang := number%5 == 0
	addPlong := number%7 == 0

	if !addPling && !addPlang && !addPlong {
		return strconv.Itoa(number)
	}

	var result string

	if addPling {
		result += "Pling"
	}
	if addPlang {
		result += "Plang"
	}
	if addPlong {
		result += "Plong"
	}

	return result
}
