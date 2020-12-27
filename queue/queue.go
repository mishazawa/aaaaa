package queue

import "github.com/mishazawa/aaaaa/deque"

type Queue struct {
	d deque.Deque
}

func NewQueue() *Queue {
	return new(Queue)
}

func (s *Queue) Enqueue(val interface{}) {
	s.d.EnqueueTail(val)
}

func (s *Queue) Dequeue() (error, interface{}) {
	return s.d.DequeueHead()
}

func (s *Queue) Peek() (error, interface{}) {
	return s.d.PeekHead()
}
