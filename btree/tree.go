package btree

import "errors"

var (
	//Threshold threshold for one node childs count
	Threshold = 256
)

//node type
const (
	FATHER uint = iota
	LEFT
)

//error type
var (
	ErrorDuplicateKkey error = errors.New("duplicate key value")
)

//Node root node
type Node struct {
	Key        uint32
	Value      string
	ChildCount int
	NodeType   uint
	LeftNode   *Node
	Pre        *Node
	Next       *Node
	Father     *Node
}

//NewTree get node tree
func NewTree(key uint32, value string) *Node {
	rootNode := NewFatherNode(key)
	leafNode := NewLeafNode(key, value)
	leafNode.Father = rootNode
	rootNode.LeftNode = leafNode
	rootNode.ChildCount = 1
	return rootNode
}

//NewLeafNode get node leafNode
func NewLeafNode(key uint32, value string) *Node {
	return &Node{
		NodeType: LEFT,
		Key:      key,
		Value:    value,
	}
}

//NewFatherNode get fahter node
func NewFatherNode(key uint32) *Node {
	return &Node{
		NodeType: FATHER,
		Key:      key,
	}
}

//GetFather get node father
func (n *Node) GetFather() *Node {
	return n.Father
}

//GetType get node type
func (n *Node) GetType() uint {
	return n.NodeType
}

//GetChildCount get child node count
func (n *Node) GetChildCount() int {
	return n.ChildCount
}

//IsRoot  node is root node
func (n *Node) IsRoot() bool {
	if n.Father == nil {
		return true
	}
	return false
}

//ChildIsLeafNode child is LeftNode
func (n *Node) ChildIsLeafNode() bool {
	if n.ChildCount != 0 && n.LeftNode.NodeType == LEFT {
		return true
	}
	return false
}

//GetNext get next node
func (n *Node) GetNext() *Node {
	return n.Next
}

//GetPre get pre node
func (n *Node) GetPre() *Node {
	return n.Pre
}

//GetCenterChildNode get center chlid node
func (n *Node) getCenterChildNode() *Node {
	centerNode := n.LeftNode
	centerCount := n.ChildCount/2 + 1
	for i := 0; i < centerCount; i++ {
		centerNode = centerNode.Next
	}
	return centerNode
}

//InsertNodeValue insert data to tree
func (n *Node) InsertNodeValue(key uint32, value string) error {
	node := NewLeafNode(key, value)
	return n.InsertNode(node)
}

//InsertNode Insert child node
func (n *Node) InsertNode(childNode *Node) (err error) {
	if childNode.Key > n.Key {
		nextNode := n.LeftNode
		i := 0
		for ; ; i++ {
			if childNode.Key < nextNode.Key {
				if n.ChildIsLeafNode() {
					preNode := nextNode.Pre
					preNode.Next = childNode
					childNode.Pre = preNode
					childNode.Next = nextNode
					nextNode.Pre = childNode
					n.ChildCount++
					n.splitChilds()
				} else {
					nextNode.Pre.InsertNode(childNode)
				}
				break
			} else if childNode.Key == nextNode.Key {
				return ErrorDuplicateKkey
			} else {
				if i == n.ChildCount-1 {
					if nextNode.NodeType == LEFT {
						childNode.Pre = nextNode
						nextNode.Next = childNode
						if nextNode.Next != nil {
							nextNode.Next.Pre = childNode
						}
						n.ChildCount++
						n.splitChilds()
					} else {
						nextNode.InsertNode(childNode)
					}
					break
				}
			}
			nextNode = nextNode.Next
		}
		return
	} else if childNode.Key == n.Key {
		return ErrorDuplicateKkey
	} else {
		return n.addInLeft(childNode)
	}

}
func (n *Node) addInLeft(childNode *Node) (err error) {
	if n.NodeType == LEFT {
		n.Father.LeftNode = childNode
		childNode.Father = n.Father
		childNode.Next = n
		n.Pre = childNode
		n.Father.ChildCount++
		n.Father.splitChilds()
		return
	}
	n.Key = childNode.Key
	return n.LeftNode.addInLeft(childNode)

}
func (n *Node) splitChilds() {
	if n.ChildCount == Threshold {
		center := Threshold/2 + 1
		centerNode := n.LeftNode
		for i := 1; i < center; i++ {
			centerNode = centerNode.Next
		}
		newFatherNode := initFatherNode(centerNode)
		rightNode := centerNode
		changeFather(rightNode, newFatherNode, Threshold-center+1)
		newFatherNode.ChildCount = Threshold - center + 1
		if n.Father != nil { //中间层拆分
			n.ChildCount = center - 1
			newFatherNode.Next = n.Next
			newFatherNode.Pre = n
			if n.Next != nil {
				n.Next.Pre = newFatherNode
			}
			n.Next = newFatherNode
			newFatherNode.Father = n.Father
			n.Father.ChildCount++
			n.Father.splitChilds()
		} else {
			leftNode := NewFatherNode(n.Key)
			leftNode.LeftNode = n.LeftNode
			cn := n.LeftNode
			changeFather(cn, leftNode, center-1)
			leftNode.ChildCount = center - 1
			leftNode.Father = n
			n.LeftNode = leftNode
			leftNode.Next = newFatherNode
			newFatherNode.Pre = leftNode
			newFatherNode.Father = n
			n.ChildCount = 2
		}
	}
}
func initFatherNode(node *Node) *Node {
	return &Node{
		Key:      node.Key,
		NodeType: FATHER,
		LeftNode: node,
	}
}

//SearchKey get key node
func (n *Node) SearchKey(key uint32) *Node {
	if key < n.Key {
		return nil
	} else if key == n.Key {
		if n.NodeType == LEFT {
			return n
		}
		return n.LeftNode.SearchKey(key)
	} else {
		nowNode := n.LeftNode
		for i := 1; i < n.ChildCount; i++ {
			nowNode = nowNode.Next
			if key == nowNode.Key {
				if nowNode.NodeType == LEFT {
					return nowNode
				}
				return nowNode.SearchKey(key)

			} else if key < nowNode.Key {
				return nowNode.Pre.SearchKey(key)
			} else if key > nowNode.Key && nowNode.Next == nil && nowNode.NodeType == FATHER {
				return nowNode.SearchKey(key)
			}
		}
		return nil
	}
}
func changeFather(fristNode *Node, fahterNode *Node, nodeCount int) {
	for i := 0; i < nodeCount; i++ {
		fristNode.Father = fahterNode
		fristNode = fristNode.Next
	}
}
