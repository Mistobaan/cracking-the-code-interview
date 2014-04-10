package chapter1

import (
	"sort"
)

type RuneSlice []rune

func (rs RuneSlice) Len() int           { return len(rs) }
func (rs RuneSlice) Less(i, j int) bool { return rs[i] < rs[j] }
func (rs RuneSlice) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }

func SortString(a string) string {
	result := make([]rune, len(a))
	n := 0
	for _, r := range result {
		result[n] = r
		n += 1
	}

	result = result[0:n]
	sort.Sort(RuneSlice(result))

	return string(result)
}

func IsPermutationV1(a, b string) bool {
	return SortString(a) == SortString(b)
}

type RuneCounter struct {
	m map[rune]int
}

func NewRuneCounter() *RuneCounter {
	return &RuneCounter{
		m: make(map[rune]int),
	}
}

func (r *RuneCounter) Keys() []rune {
	var keys []rune
	for k := range r.m {
		keys = append(keys, k)
	}
	return keys
}

func (rc *RuneCounter) Inc(key rune) (result int) {
	if value, ok := rc.m[key]; !ok {
		result = 1
	} else {
		result = value + 1
	}
	rc.m[key] = result
	return result
}

func (r *RuneCounter) Equals(other RuneCounter) bool {
	k1 := r.Keys()
	k2 := r.Keys()

	if len(k1) != len(k2) {
		return false
	}

	for _, key := range k1 {
		v := r.m[key]

		if value, ok := other.m[key]; !ok {
			return false
		} else if value != v {
			return false
		}
	}

	return true
}

func CountRunes(s string) *RuneCounter {
	rc := NewRuneCounter()
	for _, key := range s {
		rc.Inc(key)
	}
	return rc
}

func IsPermutationV2(a, b string) bool {
	return CountRunes(a) == CountRunes(b)
}
