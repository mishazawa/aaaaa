package list

import "testing"

func TestNodeCreate(t *testing.T) {
	value := 1
	head := NewNode(value)

	if head.Next != nil {
		t.Errorf("Node.Next = %v; want nil", head.Next)
	}

	if head.Prev != nil {
		t.Errorf("Node.Prev = %v; want nil", head.Prev)
	}

	if head.Value != value {
		t.Errorf("Node.Value = %v; want %v", head.Value, value)
	}
}

func TestLinkedListPrepend(t *testing.T) {
	list := NewLinkedList()
	testval1 := "first value"
	testval2 := "second value"
	testval3 := "third value"

	if list.Len() != 0 {
		t.Errorf("LinkedList.Len = %v; want %v", list.Len(), 0)
	}

	list.Prepend(testval1)

	if list.Len() != 1 {
		t.Errorf("LinkedList.Len = %v; want %v", list.Len(), 1)
	}

	head := list.Head()
	tail := list.Tail()

	if head.Value != testval1 {
		t.Errorf("LinkedList.Head.Value = %v; want %v", head.Value, testval1)
	}

	if tail.Value != testval1 {
		t.Errorf("LinkedList.Tail.Value = %v; want %v", tail.Value, testval1)
	}

	list.Prepend(testval2)

	if list.Len() != 2 {
		t.Errorf("LinkedList.Len = %v; want %v", list.Len(), 2)
	}

	head = list.Head()
	tail = list.Tail()

	if tail.Value != testval1 {
		t.Errorf("LinkedList.Head.Value = %v; want %v", tail.Value, testval1)
	}

	if head.Value != testval2 {
		t.Errorf("LinkedList.Tail.Value = %v; want %v", head.Value, testval2)
	}

	if tail.Prev != head {
		t.Errorf("Node.Prev = %v; want %v", tail.Prev, head)
	}

	if head.Next != tail {
		t.Errorf("Node.Next = %v; want %v", head.Next, tail)
	}

	list.Prepend(testval3)

	head = list.Head()
	tail = list.Tail()

	if tail.Value != testval1 {
		t.Errorf("LinkedList.Head.Value = %v; want %v", tail.Value, testval1)
	}

	if head.Value != testval3 {
		t.Errorf("LinkedList.Tail.Value = %v; want %v", head.Value, testval3)
	}

	if head.Next != tail.Prev {
		t.Errorf("Node.Prev = %v; want %v", head.Prev, tail.Prev)
	}
}

func TestLinkedListAppend(t *testing.T) {
	list := NewLinkedList()
	testval1 := "first value"
	testval2 := "second value"
	testval3 := "third value"

	list.Append(testval1)
	list.Append(testval2)
	list.Append(testval3)

	head := list.Head()
	tail := list.Tail()

	if head.Value != testval1 {
		t.Errorf("LinkedList.Head.Value = %v; want %v", head.Value, testval1)
	}

	if tail.Value != testval3 {
		t.Errorf("LinkedList.Tail.Value = %v; want %v", tail.Value, testval3)
	}

	if head.Next != tail.Prev {
		t.Errorf("Node.Next = %v; want %v", head.Next, tail.Prev)
	}
}

func TestLinkedListFind(t *testing.T) {
	list := NewLinkedList()

	testval1 := 1
	testval2 := 2
	testval3 := 3

	list.Append(testval1)
	list.Append(testval2)
	list.Append(testval3)

	resultOk := list.Find(testval3)
	resultNotFound := list.Find(666)

	if resultOk == nil {
		t.Errorf("LinkedList.Find testval3 should be found")
	}

	if resultOk != nil && resultOk.Value != testval3 {
		t.Errorf("LinkedList.Find testval3 = %v; want %v", resultOk, testval3)
	}

	if resultNotFound != nil {
		t.Errorf("LinkedList.Find 666 should not be found; have = %v", resultNotFound)
	}
}

func TestLinkedListContains(t *testing.T) {
	list := NewLinkedList()

	testval1 := 1
	testval2 := 2
	testval3 := 3

	list.Append(testval1)
	list.Append(testval2)
	list.Append(testval3)

	resultOk := list.Contains(testval3)
	resultNotFound := list.Contains(666)

	if resultOk != true {
		t.Errorf("LinkedList.Find testval3 should be found")
	}

	if resultNotFound != false {
		t.Errorf("LinkedList.Find 666 should not be found")
	}
}

func TestLinkedListRemove(t *testing.T) {
	list := NewLinkedList()

	testval1 := 1
	testval2 := 2
	testval3 := 3

	list.Append(testval1)
	list.Append(testval2)
	list.Append(testval2)
	list.Append(testval3)

	initLen := list.Len()

	list.Remove(testval2)

	if list.Len() != initLen-1 {
		t.Errorf("LinkedList.Len = %v; want %v", list.Len(), initLen-1)
	}

	initLen = list.Len()

	testnode2 := list.Find(testval2)

	if testnode2.Next.Value != testval3 {
		t.Errorf("node2.Next = %v; want %v", testnode2.Next.Value, testval3)
	}

	list.Remove(testval2)
	initLen = list.Len()

	testnode1 := list.Find(testval1)
	testnode3 := list.Find(testval3)

	if testnode1.Next != testnode3 {
		t.Errorf("node1.Next = %v; want %v", testnode1.Next, testnode3)
	}

	if testnode3.Prev != testnode1 {
		t.Errorf("node3.Prev = %v; want %v", testnode3.Prev, testnode1)
	}

	list.Remove(testval1)
	testnode3 = list.Find(testval3)

	if testnode3.Prev != nil {
		t.Errorf("node3.Prev = %v; want nil", testnode3.Prev)
	}

	initLen = list.Len()

	list.Remove(testval1)

	if list.Len() != initLen {
		t.Errorf("LinkedList.Len = %v; want %v", list.Len(), initLen)
	}

	if list.Head() != testnode3 {
		t.Errorf("LinkedList.Head = %v; want %v", list.Head(), testnode3)
	}

	if list.Tail() != testnode3 {
		t.Errorf("LinkedList.Tail = %v; want %v", list.Tail(), testnode3)
	}
}

func TestLinkedListRange(t *testing.T) {
	list := NewLinkedList()

	testval1 := 1
	testval2 := 2
	testval3 := 3

	list.Append(testval1)
	list.Append(testval2)
	list.Append(testval3)

	buffer := make([]int, 0)
	list.Range(func(val interface{}) {
		i := val.(int)
		buffer = append(buffer, i)
	})

	if buffer[0] != testval1 || buffer[1] != testval2 || buffer[2] != testval3 {
		t.Errorf("buffer[] = %v; want %v", buffer, []int{testval1, testval2, testval3})
	}
}

func TestLinkedListRangeBackward(t *testing.T) {
	list := NewLinkedList()

	testval1 := 1
	testval2 := 2
	testval3 := 3

	list.Append(testval1)
	list.Append(testval2)
	list.Append(testval3)

	buffer := make([]int, 0)
	list.RangeBackward(func(val interface{}) {
		i := val.(int)
		buffer = append(buffer, i)
	})

	if buffer[0] != testval3 || buffer[1] != testval2 || buffer[2] != testval1 {
		t.Errorf("buffer[] = %v; want %v", buffer, []int{testval3, testval2, testval1})
	}
}
