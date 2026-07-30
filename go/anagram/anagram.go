package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	lengthFilteredCanidates := filterOutLengthMatches(subject, candidates)

	for i, element := range lengthFilteredCanidates {
		lengthFilteredCanidates[i] = strings.ToLower(element)
	}

	var uniqueCandidates []string
	for _, element := range lengthFilteredCanidates {
		if element == subject {
			continue
		}
		uniqueCandidates = append(uniqueCandidates, element)
	}

	panic("Please implement the Detect function")
}

func filterOutLengthMatches(subject string, candidates []string) []string {
	subjectLength := len(subject)
	var possibleCanidates []string

	for _, element := range candidates {
		if len(element) != subjectLength {
			continue
		}

		possibleCanidates = append(possibleCanidates, element)
	}

	return possibleCanidates
}
