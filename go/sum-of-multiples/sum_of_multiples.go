package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var sum int

	for i := 1; i < limit; i++ {
		if isMultipleOfDivisors(i, divisors) {
			sum += i
		}
	}

	return sum
}

func isMultipleOfDivisors(number int, divisors []int) bool {
	for _, divisor := range divisors {
		if divisor != 0 && number%divisor == 0 {
			return true
		}
	}
	return false
}
