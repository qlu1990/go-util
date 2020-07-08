package btree

import (
	"fmt"
	"strconv"
	"testing"
)

var (
	b = NewTree(1, "abc")
)

func TestInsert(t *testing.T) {
	for i := 0; i < 1000; i++ {
		value := strconv.Itoa(i)
		b.InsertNodeValue(uint32(i), value)
	}
}

func TestSearchData(t *testing.T) {
	n := b.SearchKey(100)
	next := n
	for {
		if next == nil {
			break
		}
		fmt.Println(next.Key, next.Value)
		next = next.Next
	}
}
