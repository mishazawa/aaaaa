package list

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

func NewNode(val interface{}) *Node {
	return &Node{Value: val, Next: nil, Prev: nil}
}

type LinkedList struct {
	count uint
	head  *Node
	tail  *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{0, nil, nil}
}

func (list *LinkedList) Len() uint {
	return list.count
}

func (list *LinkedList) Head() *Node {
	return list.head
}

func (list *LinkedList) Tail() *Node {
	return list.tail
}

func (list *LinkedList) Prepend(val interface{}) {
	node := NewNode(val)
	node.Next = list.head

	if list.head != nil {
		list.head.Prev = node
	}

	list.head = node

	if list.tail == nil {
		list.tail = node
	}

	list.count += 1
}

func (list *LinkedList) Append(val interface{}) {
	if list.tail == nil {
		list.Prepend(val)
	} else {
		node := NewNode(val)

		list.tail.Next = node
		node.Prev = list.tail

		list.tail = node
		list.count += 1
	}
}

func (list *LinkedList) Find(val interface{}) *Node {
	curr := list.head
	for curr != nil {
		if curr.Value == val {
			return curr
		}
		curr = curr.Next
	}
	return nil
}

func (list *LinkedList) Contains(val interface{}) bool {
	return list.Find(val) != nil
}
