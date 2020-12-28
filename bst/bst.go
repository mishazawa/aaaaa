package bst

type ComparableValue interface {
	GT(interface{}) bool
}

type Node struct {
	Value ComparableValue
	Left  *Node
	Right *Node
}

type actionFn func(interface{})

func NewRoot(val ComparableValue) *Node {
	return &Node{val, nil, nil}
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
