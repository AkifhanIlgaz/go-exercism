package lsproduct

import (
	"errors"
	"strconv"
)

var (
	ErrNegativeSpan = errors.New("span must not be negative")
	ErrInvalidDigit = errors.New("digits input must only contain digits")
	ErrBigSpan      = errors.New("span must be smaller than string length")
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}

	if len(digits) < span {
		return 0, ErrBigSpan
	}

	var maxProduct int64

	for i := 0; i < len(digits)-span+1; i++ {
		consecutive := digits[i : i+span]
		currentProduct, err := calcProduct(consecutive)
		if err != nil {
			return 0, ErrInvalidDigit
		}
		maxProduct = Max(maxProduct, currentProduct)
	}

	return maxProduct, nil
}

func calcProduct(digits string) (int64, error) {

	product := 1

	for _, d := range digits {
		num, err := strconv.Atoi(string(d))
		if err != nil {
			return 0, ErrInvalidDigit
		}
		product *= num
	}

	return int64(product), nil
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
