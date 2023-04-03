package diffsquares

func SquareOfSum(n int) int {
	var sum int
	for n > 0 {
		sum += n
		n--
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	var sum int
	for n > 0 {
		sum += n * n
		n--
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
