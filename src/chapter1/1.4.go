package chapter1

import (
	"strings"
	"unicode"
)

func ReplaceWhiteSpaceV1(s string) string {
	s = strings.TrimSpace(s)
	return strings.Join(strings.Split(s, " "), "%20")
}

func ReplaceWhiteSpaceV2(s string) string {
	spaces := 0
	n := 0
	trimLeft := 0  // first non space char
	trimRight := 0 // Last non space char

	for _, r := range s {
		if unicode.IsSpace(r) {
			if n == 0 {
				continue
			}
			spaces += 1
		} else {
			if n == 0 {
				trimLeft = n
			}

			trimRight = n
		}

		n += 1
	}

	result := make([]rune, (trimRight-trimLeft)+(spaces*2))

	{
		n := 0
		for _, r := range s[trimLeft:trimRight] {
			if unicode.IsSpace(r) {
				result[n] = '%'
				result[n+1] = '2'
				result[n+2] = '0'
				n += 3
				continue
			}
			result[n] = r
			n += 1
		}

	}

	return string(result[0:n])
}
