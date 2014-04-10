package chapter1

import (
	"testing"
)

func BenchmarkIsPermutationV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPermutationV1("123456789ABCDE", "ABCDE123456789")
	}
}

func BenchmarkIsPermutationV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPermutationV2("123456789ABCDE", "ABCDE123456789")
	}
}
