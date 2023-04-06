// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind = NaT

	if IsTriAngle(a, b, c) {
		switch {
		case a == b && b == c:
			k = Equ
		case a == b, a == c, b == c:
			k = Iso
		default:
			k = Sca
		}
	}

	return k
}

func IsTriAngle(lengths ...float64) bool {
	max, total := 0., 0.

	for _, len := range lengths {
		if len <= 0 {
			return false
		}
		max = math.Max(max, len)
		total += len
	}

	return total-max >= max
}
