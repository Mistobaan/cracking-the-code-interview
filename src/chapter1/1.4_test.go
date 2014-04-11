package chapter1

import (
	"testing"
)

func BenchmarkReplaceWhiteSpaceV1(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		ReplaceWhiteSpaceV1("          MR John Smith               ")
	}
}

func BenchmarkReplaceWhiteSpaceV2(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		ReplaceWhiteSpaceV2("          MR John Smith               ")
	}
}
