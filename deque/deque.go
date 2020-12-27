package deque

import (
	"errors"
	"github.com/mishazawa/aaaaa/list"
)

type Deque struct {
	list list.LinkedList
}

func NewDeque() *Deque {
	return new(Deque)
}

func (d *Deque) Len() int {
	return d.list.Len()
}

func (d *Deque) DequeueHead() (error, interface{}) {
	err, val := d.PeekHead()

	if err == nil {
		d.list.RemoveHead()
	}

	return err, val
}

func (d *Deque) DequeueTail() (error, interface{}) {
	err, val := d.PeekTail()

	if err == nil {
		d.list.RemoveTail()
	}

	return err, val
}

func (d *Deque) EnqueueHead(val interface{}) {
	d.list.Prepend(val)
}

func (d *Deque) EnqueueTail(val interface{}) {
	d.list.Append(val)
}

func (d *Deque) PeekHead() (error, interface{}) {
	node := d.list.Head()

	if node != nil {
		return nil, node.Value
	}

	return errors.New("head is empty"), node
}

func (d *Deque) PeekTail() (error, interface{}) {
	node := d.list.Tail()

	if node != nil {
		return nil, node.Value
	}

	return errors.New("tail is empty"), node
}
