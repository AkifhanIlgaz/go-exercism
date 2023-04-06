package grains

import "fmt"

// There are 64 squares on a chessboard (where square 1 has one grain, square 2 has two grains, and so on).
// how many grains were on a given square, and
func Square(number int) (uint64, error) {
	num := uint64(number)
	if num < 1 || num > 64 {
		return 0, fmt.Errorf("invalid square number: %d", number)
	}

	// 1 << (num - 1) == 2 ^ (num - 1)
	return 1 << (num - 1), nil
}

func Total() uint64 {
	return 1<<64 - 1
}
