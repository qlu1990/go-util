package main

import (
	"fmt"
	"strconv"

	"github.com/qlu1990/go-util/btree"
)

func main() {
	tree := btree.NewTree(1001, "aaa")
	for i := 1; i < 10000; i++ {
		tree.InsertNodeValue(uint32(i), strconv.Itoa(i))
	}
	for i := 2; i < 1002; i++ {
		node := tree.SearchKey(uint32(i))
		if node != nil {
			fmt.Printf("key: %d,value %s\n", i, node.Value)
		}
	}
	// node := tree.SearchKey(2)
	// for i := 1; i < 1000; i++ {
	// 	fmt.Printf("key: %d,value %s\n", node.Key, node.Value)
	// 	node = node.Next
	// }

}
