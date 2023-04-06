package resistorcolorduo

import (
	"fmt"
	"strconv"
)

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {

	var val string

	for _, color := range colors[:2] {
		val += ColorCode(color)
	}

	numVal, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println(err)
	}

	return numVal
}

func ColorCode(color string) string {
	var colorcode string

	switch color {
	case "black":
		colorcode = "0"
	case "brown":
		colorcode = "1"
	case "red":
		colorcode = "2"
	case "orange":
		colorcode = "3"
	case "yellow":
		colorcode = "4"
	case "green":
		colorcode = "5"
	case "blue":
		colorcode = "6"
	case "violet":
		colorcode = "7"
	case "grey":
		colorcode = "8"
	case "white":
		colorcode = "9"
	}

	return colorcode
}
