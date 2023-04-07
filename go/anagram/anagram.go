package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)

	validAnagrams := []string{}

	for _, candidate := range candidates {
		if isAnagram(subject, strings.ToLower(candidate)) {
			validAnagrams = append(validAnagrams, candidate)
		}
	}

	return validAnagrams
}

func isAnagram(subject string, candidate string) bool {
	if candidate == subject || len(candidate) != len(subject) {
		return false
	}

	for _, chr := range subject {
		if strings.Count(subject, string(chr)) != strings.Count(candidate, string(chr)) {
			return false
		}
	}
	return true
}
