package binary

import (
	"testing"
)

func Test_Tree_Insert(t *testing.T) {

	var tree Tree

	for i := 0; i < 100; i += 1 {
		tree.Insert(&Node{Key: i})
	}

	for i := 0; i < 100; i += 1 {
		tree.Delete(tree.Search(i))
	}

}
