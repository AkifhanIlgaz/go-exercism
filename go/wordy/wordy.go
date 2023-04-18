package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	question, validQuestion := strings.CutPrefix(question, "What is ")
	if !validQuestion {
		return 0, false
	}

	question, _ = strings.CutSuffix(question, "?")

	tokens := strings.Fields(question)

	if len(tokens) == 1 {
		result, err := strconv.Atoi(tokens[0])
		if err != nil {
			return 0, false
		}
		return result, true
	}

	var result int

	var operands []int

	for i := 0; i < len(tokens); i += 2 {
		if tokens[i] == "by" {
			i--
			continue
		}
		operand, err := strconv.Atoi(tokens[i])
		if err != nil {
			return 0, false
		}
		operands = append(operands, operand)
	}

	if len(operands) == 1 {
		return 0, false
	}

	for i := 1; i < len(tokens); i += 2 {

		switch tokens[i] {
		case "plus":
			result = Add(operands[0], operands[1])
		case "minus":
			result = Subtract(operands[0], operands[1])
		case "multiplied":
			i++
			result = Multiply(operands[0], operands[1])
		case "divided":
			i++
			result = Divide(operands[0], operands[1])
		default:
			return 0, false
		}

		operands = append([]int{result}, operands[2:]...)
	}

	return result, true
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	return a / b
}
