package btree

import (
	"errors"
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

//Node tree  node
type Node struct {
	key      string
	value    string
	nodeType uint8
	father   *Node
	childs   []*Node
	next     *Node
}

//NewNode create a node
func NewNode(key string, value string) *Node {
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
func (t *Node) Insert(n *Node) (err error) {
	if t.nodeType == LEAF {
		childsLen := len(t.childs)
		n.father = t
		if childsLen == 0 {
			t.childs = append(t.childs, n)
			return
		}
		poi := 0
		for i, cn := range t.childs {
			poi = i
			if strings.Compare(n.key, cn.key) < 0 {
				break
			}

		}

		if strings.Compare(n.key, t.childs[poi].key) < 0 { /*   */
			t.childs = append(t.childs, n)
			copy(t.childs[poi+1:], t.childs[poi:])
			t.childs[poi+1] = n
			n.next = t.childs[poi].next
			t.childs[poi].next = n
			n.key, n.value, t.childs[poi].key, t.childs[poi].value =
				t.childs[poi].key, t.childs[poi].value, n.key, n.value
		} else if strings.Compare(n.key, t.childs[poi].key) == 0 {
			return errors.New("insert repeat key")
		} else {
			t.childs = append(t.childs, n)
			n.next = t.childs[childsLen-1].next
			t.childs[childsLen-1].next = n
		}

		if len(t.childs) == Threshold {
			t.nodeSplit()
		}
		return
	}

	childsLen := len(t.childs)
	if childsLen == 0 {
		NewNode := NewNode(n.key, n.value)
		NewNode.nodeType = LEAF
		t.childs = append(t.childs, NewNode)
		n.father = NewNode
		NewNode.Insert(n)
		return
	}
	index := 0

	for i, cn := range t.childs {
		index = i
		if index == childsLen-1 || (strings.Compare(n.key, cn.key) >= 0 &&
			strings.Compare(n.key, t.childs[i+1].key) < 0) {
			break
		}
	}

	if index > 0 {
		t.childs[index].Insert(n)
	} else {
		if strings.Compare(n.key, t.childs[0].key) < 0 {
			leafNode := NewNode(n.key, n.value)
			leafNode.nodeType = LEAF
			n.father = leafNode
			leafNode.childs = append(leafNode.childs, n)
			n.next = t.childs[index].childs[0]
			leafNode.father = t
			t.childs = append(t.childs, leafNode)
			copy(t.childs[1:], t.childs)
			t.childs[0] = leafNode
			if len(t.childs) == Threshold {
				t.nodeSplit()
			}
		} else {
			t.childs[0].Insert(n)
		}

	}
	return

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
			node1 := NewNode(t.childs[centerPoi].key, t.childs[centerPoi].key)
			node1.nodeType = nodetype
			node1.childs = make([]*Node, centerPoi+1)
			copy(node1.childs, t.childs[:centerPoi+1])
			node2 := NewNode(t.childs[Threshold-1].key, t.childs[Threshold-1].key)
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
		newLeafNode := NewNode(centerNode.key, centerNode.value)
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

//SearchData search  data from tree
func (t *Node) SearchData(data string) *Node {
	switch t.nodeType {
	case LEAF:
		{
			for _, cn := range t.childs {
				if strings.Compare(data, cn.key) <= 0 {
					return cn
				}
			}
			return nil
		}
	default:
		{
			childsLen := len(t.childs)
			for i, cn := range t.childs {
				if strings.Compare(data, cn.key) < 0 {
					return nil
				}
				if i == childsLen-1 || strings.Compare(data, t.childs[i+1].key) < 0 {
					return cn.SearchData(data)
				}
			}
			return nil
		}
	}

}

//GetValue get node value
func (t *Node) GetValue() string {
	return t.value
}

//GetNext get next node point
func (t *Node) GetNext() *Node {
	return t.next
}
