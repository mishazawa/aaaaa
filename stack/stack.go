package stack

import "github.com/mishazawa/aaaaa/deque"

type Stack struct {
	d deque.Deque
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Push(val interface{}) {
	s.d.EnqueueTail(val)
}

func (s *Stack) Pop() (error, interface{}) {
	return s.d.DequeueTail()
}

func (s *Stack) Peek() (error, interface{}) {
	return s.d.PeekTail()
}
