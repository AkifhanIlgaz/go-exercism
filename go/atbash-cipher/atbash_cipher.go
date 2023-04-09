package atbash

import (
	"strings"
	"unicode"
)

const cipher = "zyxwvutsrqponmlkjihgfedcba"
const chunkSize = 5

func Atbash(s string) string {
	chunks := []string{}

	letters := strings.FieldsFunc(s, func(r rune) bool { return unicode.IsPunct(r) || unicode.IsSpace(r) })
	s = strings.Join(letters, "")

	var chunk strings.Builder

	for _, chr := range s {
		chr = unicode.ToLower(chr)

		if chunk.Len() == chunkSize {
			chunks = append(chunks, chunk.String())
			chunk.Reset()
		}

		writeToChunk(&chunk, chr)

	}

	chunks = append(chunks, chunk.String())

	return strings.Join(chunks, " ")
}

func writeToChunk(sb *strings.Builder, chr rune) {
	if unicode.IsDigit(chr) {
		sb.WriteRune(chr)
	} else if unicode.IsLetter(chr) {
		sb.WriteByte(cipher[chr-'a'])
	}
}
