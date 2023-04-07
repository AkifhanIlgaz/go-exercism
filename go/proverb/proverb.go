package proverb

import "fmt"

const sentenceOne = "For want of a %s the %s was lost."
const sentenceLast = "And all for the want of a %s."

func Proverb(rhyme []string) []string {
	p := make([]string, len(rhyme))
	for i := range rhyme {
		if i < len(rhyme)-1 {
			p[i] = fmt.Sprintf(sentenceOne, rhyme[i], rhyme[i+1])
		} else {
			p[i] = fmt.Sprintf(sentenceLast, rhyme[0])
		}
	}
	return p
}
