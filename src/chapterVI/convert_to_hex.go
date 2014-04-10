package chapterVI

import (
	"math"
)

func digit(char byte) int {
	if char >= '0' && char <= '9' {
		return int(char - '0')
	}
	if char >= 'A' && char <= 'F' {
		return int(char-'A') + 10
	}

	if char >= 'a' && char <= 'f' {
		return int(char-'a') + 10
	}
	return -1
}

func convertToDigit(input string, base int) int {
	result := 0
	if len(input) == 0 || base < 2 {
		return -1
	}
	for i := len(input) - 1; i >= 0; i += -1 {
		ch := input[i]
		value := digit(ch)
		if value > base {
			return -1
		}
		position := len(input) - 1 - i
		result += value * int(math.Pow(float64(base), float64(position)))
	}

	return result
}

func match(binary, hexa string) bool {
	a := convertToDigit(binary, 2)
	b := convertToDigit(hexa, 16)

	if a == -1 || b == -1 {
		return false
	}

	return a == b
}
