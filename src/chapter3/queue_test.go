package chapter3

import (
	"testing"
)

func Test_QueuePush_MustWork(t *testing.T) {
	q := NewQueue()
	q.Push(0)
	q.Pop()
	if _, err := q.Pop(); err == nil {
		t.Fatalf("Should have failed")
	}

	for i := 0; i < 2000; i += 1 {
		q.Push(0)
	}

	for i := 0; i < 2000; i += 1 {
		q.Pop()
	}

	for i := 0; i < 4000; i += 1 {
		q.Push(0)
	}

}

func Benchmark_QueuePush(b *testing.B) {
	q := NewQueue()
	for i := 0; i < b.N; i += 1 {
		q.Push(i)
	}
}
