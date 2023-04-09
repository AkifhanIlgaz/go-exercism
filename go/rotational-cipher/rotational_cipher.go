package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	var cipheredText strings.Builder

	for _, char := range plain {
		if unicode.IsLetter(char) {
			cipheredText.WriteRune(shift(char, shiftKey))
		} else {
			cipheredText.WriteRune(char)
		}
	}

	return cipheredText.String()

}

func shift(char rune, shiftKey int) rune {
	base := 'a'

	if unicode.IsUpper(char) {
		base = 'A'
	}

	return (char-base+rune(shiftKey))%26 + base
}
