package chapter1

import (
	"sort"
)

type StringSlice []uint8

func (s StringSlice) Len() int           { return len(s) }
func (s StringSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s StringSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func HasUniqueChar(a string) bool {
	test := make(map[uint8]bool)
	if len(a) < 2 {
		return true
	}
	for i := 0; i < len(a); i += 1 {
		if _, ok := test[a[i]]; ok {
			return false
		}
		test[a[i]] = true
	}
	return true
}

func HasUniqueCharV1(a string) bool {
	if len(a) < 2 {
		return true
	}

	sort.Sort(StringSlice(a))

	for i := 0; i < len(a)-1; i += 1 {
		if a[i] == a[i+1] {
			return false
		}
	}
	return true
}
