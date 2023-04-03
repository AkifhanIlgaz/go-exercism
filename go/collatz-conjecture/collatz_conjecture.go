package collatzconjecture

import "errors"

// Return the number of steps required to reach 1 using the Collatz Conjecture.
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid input")
	}

	var count int

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		count++
	}

	return count, nil
}
