package chapter1

import (
	"fmt"
)

// http://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func reverse(s string) string {
	if len(s) < 2 {
		return s
	}

	result := make([]rune, len(s))

	count := 0
	for _, r := range s {
		result[count] = r
		count += 1
	}

	result = result[0:count]

	// Reverse
	for i := 0; i < (count >> 1); i += 1 {
		result[i], result[count-i-1] = result[count-i-1], result[i]
	}

	return string(result)
}

func main() {

	for _, tc := range []string{
		"",
		"0123456789\u0041",
	} {
		fmt.Printf("%s\n", reverse(tc))
	}

}
