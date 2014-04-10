package chapter1

import (
	"testing"
)

func DoTestUniqueCharFunction(t *testing.T, f func(a string) bool) {
	type TestCase struct {
		input  string
		result bool
	}
	test_cases := []TestCase{
		{"", true},
		{"1", true},
		{"ab", true},
		{"aa", false},
	}

	for idx, tc := range test_cases {
		if f(tc.input) != tc.result {
			t.Fatalf("%d: expected %v", idx, tc.result)
		}
	}
}

func Test_HasUniqueChar_MustWork(t *testing.T) {
	DoTestUniqueCharFunction(t, HasUniqueChar)
}

func Test_HasUniqueCharV1_MustWork(t *testing.T) {
	DoTestUniqueCharFunction(t, HasUniqueCharV1)
}
