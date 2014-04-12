package sentinel

import (
	"testing"
)

func Test_LinkedList_Insert(t *testing.T) {
	l := NewLinkedList()

	for i := 0; i < 100; i += 1 {
		l.Insert(&Node{Key: i})
	}

	for i := 0; i < 100; i += 1 {
		l.Delete(l.Search(i))
	}

}

func Benchmark_LinkedList_Insert(b *testing.B) {
	l := NewLinkedList()

	for i := 0; i < b.N; i += 1 {
		l.Insert(&Node{Key: i})
	}
}
