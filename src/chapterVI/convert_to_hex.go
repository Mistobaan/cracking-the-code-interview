package chapterVI

// exp values: 0 255
// int32 = - 2 ** 31 to 2 ** 31
//
func ipow(base int, exp uint) int {
	switch {
	case base == 0:
		return 0
	case exp == 0:
		return 1
	}

	// http://en.wikipedia.org/wiki/Exponentiation_by_squaring
	result := 1
	for exp != 0 {
		if (exp & 1) != 0 { // is odd
			result *= base
		}
		exp >>= 1
		base *= base
	}

	return result
}

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
		if value < 0 && value > base {
			return -1
		}
		position := uint(len(input) - 1 - i)
		result += value * ipow(base, position)
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
