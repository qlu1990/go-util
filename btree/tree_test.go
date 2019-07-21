package btree

import "testing"

var (
	btree = NewTree()
)

func TestInsert(t *testing.T) {
	node1 := newNode("abc", "abc")
	btree.insert(node1)
	node2 := newNode("def", "def")
	btree.insert(node2)
}
