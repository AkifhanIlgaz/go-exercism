package darts

import "math"

func Score(x, y float64) int {
	var point int

	dist := distance(x, y)
	switch {
	case dist <= 1:
		point = 10
	case dist <= 5:
		point = 5
	case dist <= 10:
		point = 1
	}

	return point
}

func distance(x, y float64) float64 {
	return math.Sqrt((x * x) + (y * y))
}
