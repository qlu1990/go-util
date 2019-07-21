package btree

import (
	"strings"
)

//node type
const (
	ROOT uint8 = iota
	FATHER
	LEAF
	DATA
)

var (
	//Threshold threshold for one node childs count
	Threshold = 256
)

// type Tree struct {
// 	nodes []*Node
// }

//Node tree  node
type Node struct {
	key      string
	value    string
	nodeType uint8
	father   *Node
	childs   []*Node
	next     *Node
}

func newNode(key string, value string) *Node {
	return &Node{
		nodeType: DATA,
		key:      key,
		value:    value,
		childs:   make([]*Node, 0),
	}

}

//NewTree get new tree
func NewTree() *Node {
	return &Node{
		nodeType: ROOT,
		childs:   make([]*Node, 0),
	}
}

//Insert data
func (t *Node) insert(n *Node) {
	if t.nodeType == LEAF {
		if len(t.childs) == 0 {
			t.childs = append(t.childs, n)
			return
		}
		poi := 0
		for i, cn := range t.childs {
			if strings.Compare(n.key, cn.key) <= 0 {
				poi = i
			}

		}
		n.father = t
		t.childs = append(t.childs, n)
		copy(t.childs[poi+1:], t.childs[poi:])
		t.childs[poi] = n
		t.childs[poi-1].next = n
		n.next = t.childs[poi+1]
		if len(t.childs) == Threshold {
			t.nodeSplit()
		}
	} else {
		if len(t.childs) == 0 {
			newNode := newNode(n.key, n.value)
			newNode.nodeType = LEAF
			t.childs = append(t.childs, newNode)
			newNode.insert(n)
			return
		}
		index := 0
		for i, cn := range t.childs {
			if strings.Compare(n.key, cn.key) <= 0 {

				index = i
				break

			}
		}
		childsLen := len(t.childs)
		if index < childsLen-1 {
			t.childs[index].insert(n)
		} else {
			if strings.Compare(n.key, t.childs[index].key) > 0 {
				leafNode := newNode(n.key, n.value)
				leafNode.nodeType = LEAF
				n.father = leafNode
				leafNode.childs = append(leafNode.childs, n)
				n.next = t.childs[childsLen-1].next
				t.childs[childsLen-1].next = n.next
				leafNode.father = t
				t.childs = append(t.childs, leafNode)
				if len(t.childs) == Threshold {
					t.nodeSplit()
				}
			}
		}
	}

}

//nodeSplit split nodechilds if is length equel Threshold
func (t *Node) nodeSplit() {
	switch t.nodeType {
	case ROOT:
		{
			nodetype := FATHER
			if t.childs[0].childs[0].nodeType == DATA {
				nodetype = LEAF
			}
			centerPoi := Threshold / 2
			node1 := newNode(t.childs[centerPoi].key, t.childs[centerPoi].key)
			node1.nodeType = nodetype
			node1.childs = make([]*Node, centerPoi+1)
			copy(node1.childs, t.childs[:centerPoi+1])
			node2 := newNode(t.childs[Threshold-1].key, t.childs[Threshold-1].key)
			node2.nodeType = nodetype
			node2.childs = make([]*Node, Threshold-centerPoi-1)
			copy(node1.childs, t.childs[centerPoi+1:])
			t.childs = t.childs[:0]
			t.childs = append(t.childs, node1)
			t.childs = append(t.childs, node2)
		}
	default:
		centerPoi := Threshold / 2
		centerNode := t.childs[centerPoi]
		newLeafNode := newNode(centerNode.key, centerNode.value)
		newLeafNode.nodeType = t.nodeType
		newLeafNode.childs = make([]*Node, centerPoi+1)
		copy(newLeafNode.childs, t.childs[:centerPoi+1])
		t.childs = t.childs[centerPoi+1:]
		fatherNode := t.father
		poi := 0
		for i, cn := range fatherNode.childs {
			if strings.Compare(newLeafNode.key, cn.key) <= 0 {
				poi = i
				break
			}
		}
		fatherNode.childs = append(fatherNode.childs, newLeafNode)
		copy(fatherNode.childs[poi+1:], fatherNode.childs[poi:])
		fatherNode.childs[poi] = newLeafNode
		if len(fatherNode.childs) == Threshold {
			fatherNode.nodeSplit()
		}
	}
}
