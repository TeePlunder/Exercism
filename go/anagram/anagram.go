package anagram

import (
	"slices"
	"strings"
)

// Detect returns the candidates that are anagrams of subject, preserving each
// candidate's original casing. A candidate equal to the subject (ignoring case)
// is not an anagram of it.
func Detect(subject string, candidates []string) []string {
	lowerSubject := strings.ToLower(subject)
	sortedSubject := sortLetters(lowerSubject)

	anagrams := []string{}
	for _, candidate := range candidates {
		lowerCandidate := strings.ToLower(candidate)
		if lowerCandidate == lowerSubject {
			continue
		}
		if sortLetters(lowerCandidate) == sortedSubject {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

// sortLetters returns the runes of word in sorted order, so that two words are
// anagrams exactly when their sorted forms are equal.
func sortLetters(word string) string {
	letters := []rune(word)
	slices.Sort(letters)
	return string(letters)
}
