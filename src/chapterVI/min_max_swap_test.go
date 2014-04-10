package chapterVI

import (
	"reflect"
	"testing"
)

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

		if !reflect.DeepEqual(tc.input, tc.output) {
			t.Fatalf("%d: %v != %v", idx, tc.input, tc.output)
		}
	}
}
