package bst

type ComparableValue interface {
	Gte(interface{}) bool
}

type Node struct {
	Value ComparableValue
	Left  *Node
	Right *Node
}

func NewRoot(val ComparableValue) *Node {
	return &Node{val, nil, nil}
}

func (n *Node) Add(val ComparableValue) {
	if n.Value == nil || n.Value.Gte(val) {
		n.addLeft(val)
	} else {
		n.addRight(val)
	}
}

func (n *Node) addRight(val ComparableValue) {
	if n.Value == nil {
		n.Value = val
		return
	}

	if n.Right == nil {
		n.Right = NewRoot(val)
		return
	}

	n.Right.Add(val)
}

func (n *Node) addLeft(val ComparableValue) {
	if n.Left == nil {
		n.Left = NewRoot(val)
		return
	}

	n.Left.Add(val)
}
