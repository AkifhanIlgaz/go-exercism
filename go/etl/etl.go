package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	transformedLetter := map[string]int{}

	for point, letters := range in {
		for _, letter := range letters {
			transformedLetter[strings.ToLower(letter)] = point
		}
	}

	return transformedLetter
}
