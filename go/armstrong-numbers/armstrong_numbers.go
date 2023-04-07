package armstrong

import "math"

func IsNumber(n int) bool {
	var sum = n
	numOfDigit := countDigit(n)

	for n > 0 {
		lastDigit := n % 10
		sum -= int(math.Pow(float64(lastDigit), numOfDigit))
		n /= 10
	}

	return sum == 0
}

func countDigit(n int) float64 {
	return math.Trunc((math.Log10(float64(n))) + 1)
}
