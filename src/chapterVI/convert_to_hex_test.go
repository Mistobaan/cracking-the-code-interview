package chapterVI

import (
	"fmt"
	"testing"
)

func Test_digit_MustWork(t *testing.T) {
	type TestCase struct {
		input  byte
		result int
	}

	test_cases := []TestCase{
		{'0', 0},
		{'a', 10},
		{'g', -1},
	}

	for idx, tc := range test_cases {
		result := digit(tc.input)

		if result != tc.result {
			t.Fatalf("%d: %d != %d", idx, result, tc.result)
		}
	}
}

func Test_convertToDigit_MustWork(t *testing.T) {
	type TestCase struct {
		input  string
		base   int
		result int
	}
	test_cases := []TestCase{
		{"10", 2, 2},
		{"10", 10, 10},
	}

	for idx, tc := range test_cases {
		result := convertToDigit(tc.input, tc.base)
		if result != tc.result {
			t.Fatalf("%d: %d != %d", idx, result, tc.result)
		}
	}

}

func Test_match_MustWork(t *testing.T) {
	type TestCase struct {
		index  int
		binary string
		hexa   string
		result bool
	}

	test_cases := []TestCase{
		{2, "", "", false},
		{3, "2", "", false},
		{4, "0", "G", false},
		{5, "-1", "", false},
		{6, "0", "0", true},
		{7, "1", "1", true},
		{8, "10", "2", true},
		{9, "11", "3", true},
		{10, "11", "3", true},
		{11, "21", "3", false},
		{12, "21", "G3", false},
		{13, "200001", "G000003", false},
		{14, "1000", "8", true},
	}

	for _, tc := range test_cases {
		t.Logf("%v", tc)
		result := match(tc.binary, tc.hexa)
		if result != tc.result {
			t.Errorf(fmt.Sprintf("Invalid result TestCase: %d %v", tc.index, tc))
		}
	}
}
