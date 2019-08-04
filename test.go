package main

import (
	"fmt"

	"github.com/qlu1990/go-util/btree"
)

func main() {
	tree := btree.NewTree()
	node1 := btree.NewNode("a", "abc")
	tree.Insert(node1)
	node2 := btree.NewNode("b", "def")
	tree.Insert(node2)
	node3 := btree.NewNode("c", "d2f")
	tree.Insert(node3)
	node4 := btree.NewNode("ba", "d5f")
	tree.Insert(node4)
	n := tree.SearchData("abc")
	next := n
	for {
		if next == nil {
			break
		}
		fmt.Println(next.GetValue())
		next = next.GetNext()
	}
}
