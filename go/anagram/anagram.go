package anagram

import (
	"slices"
	"strings"
)

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

	anagrams := []string{}
	for _, element := range lengthFilteredCanidates {
		if isSubjectAndCanidateTheSame(subject, element) {
			anagrams = append(anagrams, element)
		}
	}

	return anagrams
}

func isSubjectAndCanidateTheSame(subject, candidate string) bool {
	var sortedSubject, sortedCandidate []string

	for _, element := range subject {
		sortedSubject = append(sortedSubject, string(element))
	}

	for _, element := range candidate {
		sortedCandidate = append(sortedCandidate, string(element))
	}

	slices.Sort(sortedSubject)
	slices.Sort(sortedCandidate)

	// everything the same

	isTheSame := true
	for i := 0; i < len(sortedSubject); i++ {
		if sortedSubject[i] != sortedCandidate[i] {
			isTheSame = false
			break
		}
	}

	return isTheSame
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
