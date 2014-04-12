package collections

import (
	"testing"
)

func Test_LinkedList_MustWork(t *testing.T) {
	l := NewLinkedList()

	for i := 0; i < 100; i += 1 {
		l.Insert(&Node{Key: i})
	}

	for i := 0; i < 100; i += 1 {
		l.Delete(l.Search(i))
	}
}
