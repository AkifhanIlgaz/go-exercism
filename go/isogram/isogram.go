package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for _, letter := range word {
		if unicode.IsLetter(letter) {
			if strings.Count(word, string(letter)) > 1 {
				return false
			}
		}
	}

	return true
}
