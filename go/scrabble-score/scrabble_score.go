package scrabble

import "unicode"

/*
Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10
*/
func letterValue(letter rune) int {
	var value int

	switch letter {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		value = 1
	case 'D', 'G':
		value = 2
	case 'B', 'C', 'M', 'P':
		value = 3
	case 'F', 'H', 'V', 'W', 'Y':
		value = 4
	case 'K':
		value = 5
	case 'J', 'X':
		value = 8
	case 'Q', 'Z':
		value = 10
	}

	return value
}

func Score(word string) int {
	var score int

	for _, letter := range word {
		letter = unicode.ToUpper(letter)
		score += letterValue(letter)
	}

	return score
}
