package bst

type ComparableValue interface {
	GT(interface{}) bool
	EQ(interface{}) bool
}

type Node struct {
	Value  ComparableValue
	Left   *Node
	Right  *Node
	parent *Node
}

type actionFn func(interface{})

func NewRoot(val ComparableValue) *Node {
	return &Node{val, nil, nil, nil}
}

func newChild(val ComparableValue, parent *Node) *Node {
	return &Node{val, nil, nil, parent}
}

func (n *Node) Add(val ComparableValue) {
	if n.Value == nil || n.Value.GT(val) {
		n.addLeft(val)
	} else {
		n.addRight(val)
	}
}

func (n *Node) PreOrderTraverse(action actionFn) {
	action(n.Value)

	if n.Left != nil {
		n.Left.PreOrderTraverse(action)
	}

	if n.Right != nil {
		n.Right.PreOrderTraverse(action)
	}
}

func (n *Node) InOrderTraverse(action actionFn) {
	if n.Left != nil {
		n.Left.InOrderTraverse(action)
	}

	action(n.Value)

	if n.Right != nil {
		n.Right.InOrderTraverse(action)
	}
}

func (n *Node) PostOrderTraverse(action actionFn) {
	if n.Left != nil {
		n.Left.PostOrderTraverse(action)
	}

	if n.Right != nil {
		n.Right.PostOrderTraverse(action)
	}

	action(n.Value)
}

func (n *Node) Search(val ComparableValue) *Node {
	if n.Value.GT(val) {
		if n.Left != nil {
			return n.Left.Search(val)
		}
	} else {
		if n.Value.EQ(val) {
			return n
		}
		if n.Right != nil {
			return n.Right.Search(val)
		}
	}
	return nil
}

func (n *Node) addRight(val ComparableValue) {
	if n.Value == nil {
		n.Value = val
		return
	}

	if n.Right == nil {
		n.Right = newChild(val, n)
		return
	}

	n.Right.Add(val)
}

func (n *Node) addLeft(val ComparableValue) {
	if n.Left == nil {
		n.Left = newChild(val, n)
		return
	}

	n.Left.Add(val)
}
