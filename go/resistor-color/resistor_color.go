package resistorcolor

// Colors should return the list of all colors.
func Colors() []string {
	var colors = []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}

	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	var colorcode int

	// lowercase case

	switch color {
	case "black":
		colorcode = 0
	case "brown":
		colorcode = 1
	case "red":
		colorcode = 2
	case "orange":
		colorcode = 3
	case "yellow":
		colorcode = 4
	case "green":
		colorcode = 5
	case "blue":
		colorcode = 6
	case "violet":
		colorcode = 7
	case "grey":
		colorcode = 8
	case "white":
		colorcode = 9
	}

	return colorcode
}
