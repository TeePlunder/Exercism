package etl

import "strings"

// old list
// 1: []string{"A", "E"}, 2: []string{"D", "G"}

// goal
// "a": 1, "d": 2, "e": 1, "g": 2

func Transform(in map[int][]string) map[string]int {
	newList := make(map[string]int)
	for letterScore, letters := range in {
		for _, letter := range letters {
			lowerCaseLetter := strings.ToLower(letter)
			newList[lowerCaseLetter] = letterScore
		}
	}

	return newList
}
