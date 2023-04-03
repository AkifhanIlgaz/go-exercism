package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("dna instances must have same length")
	}

	var hammingDifference int

	for i := range a {
		if a[i] != b[i] {
			hammingDifference++
		}
	}

	return hammingDifference, nil
}
