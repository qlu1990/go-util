package btree

import (
	"fmt"
	"testing"
)

var (
	btree = NewTree()
)

func TestInsert(t *testing.T) {
	node1 := NewNode("abc", "abc")
	btree.Insert(node1)
	node2 := NewNode("def", "def")
	btree.Insert(node2)
	node3 := NewNode("dff", "d2f")
	btree.Insert(node3)
}

func TestSearchData(t *testing.T) {
	n := btree.SearchData("abc")
	next := n
	for {
		if next == nil {
			break
		}
		fmt.Println(next.key, next.value)
		next = next.next
	}
}
