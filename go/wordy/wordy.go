package wordy

import (
	"regexp"
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

func RegexAnswer(question string) (int, bool) {
	if match, _ := regexp.MatchString(`What is -?\d+(?: (?:plus|minus|divided by|multiplied by) -?\d+)*\?`, question); !match {
		return 0, false
	}

	operatorRegex := regexp.MustCompile(`(plus|minus|divided|multiplied)`)
	operators := operatorRegex.FindAllString(question, -1)

	operandRegex := regexp.MustCompile(`-?\d+`)
	operands := operandRegex.FindAllString(question, -1)

	if len(operators) != len(operands)-1 {
		return 0, false
	}

	result, _ := strconv.Atoi(operands[0])

	for i, o := range operators {
		n, _ := strconv.Atoi(operands[i+1])
		switch o {
		case "plus":
			result += n
		case "minus":
			result -= n
		case "divided":
			result /= n
		case "multiplied":
			result *= n
		}
	}

	return result, true
}
