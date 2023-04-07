package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := Frequency{}

	phrase = strings.ToLower(phrase)
	words := strings.FieldsFunc(phrase, IsValidSeparator)

	for _, word := range words {
		if word = strings.Trim(word, "'"); len(word) > 0 {
			freq[word]++
		}
	}

	return freq
}

func IsValidSeparator(chr rune) bool {
	return !(unicode.IsDigit(chr) || unicode.IsLetter(chr) || chr == '\'')
}
