package chapterVI

import (
	"testing"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i += 1 {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func Test_min_max_swap_MustWork(t *testing.T) {
	type TestCase struct {
		input  []int
		output []int
	}

	test_cases := []TestCase{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{0, 1, 2}, []int{2, 1, 0}},
	}

	for idx, tc := range test_cases {
		t.Logf("%v", tc.input)
		swap_min_max(tc.input)

		if !equal(tc.input, tc.output) {
			t.Fatalf("%d: %v != %v", idx, tc.input, tc.output)
		}
	}
}
