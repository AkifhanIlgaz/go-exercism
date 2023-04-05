package pangram

import "unicode"

func IsPangram(input string) bool {

	alphabet := [26]bool{}

	for _, char := range input {
		if unicode.IsLetter(char) {
			char = unicode.ToLower(char)
			alphabet[char-'a'] = true
		}
	}

	for _, exists := range alphabet {
		if !exists {
			return false
		}
	}

	return true
}
