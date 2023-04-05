package reverse

func Reverse(input string) string {
	var reversedString string

	for _, letter := range input {
		reversedString = string(letter) + reversedString
	}

	return reversedString
}
