package chapter3

import (
	"testing"
)

func Test_Stack_Push_MustWork(t *testing.T) {
	stack := NewStack()

	for i := 0; i < 3000; i += 1 {
		stack.Push(&struct{}{})
	}
}

func Benchmark_Stack_Push(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i += 1 {
		stack.Push(&struct{}{})
	}
}
