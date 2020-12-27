package deque

import "testing"

func TestDequeCreate(t *testing.T) {
	v := NewDeque()

	if v.Len() != 0 {
		t.Errorf("Deque.Len = %v; want %v", v.Len(), 0)
	}
}

func TestDequeHead(t *testing.T) {
	v := NewDeque()

	len := 5

	for i := 0; i < 5; i++ {
		v.EnqueueHead(i)
	}

	if v.Len() != len {
		t.Errorf("Deque.Len = %v; want %v", v.Len(), len)
	}

	err, _ := v.DequeueHead()

	if err != nil {
		t.Errorf("Deque.DequeueHead err = %v; want %v", err, nil)
	}

	for v.Len() != 0 {
		v.DequeueHead()
	}

	err, _ = v.DequeueHead()

	if err == nil {
		t.Errorf("Deque.DequeueHead err = %v; want `no head`", err)
	}
}

func TestDequeTail(t *testing.T) {
	v := NewDeque()

	len := 5

	for i := 0; i < 5; i++ {
		v.EnqueueTail(i)
	}

	if v.Len() != len {
		t.Errorf("Deque.Len = %v; want %v", v.Len(), len)
	}

	err, _ := v.DequeueTail()

	if err != nil {
		t.Errorf("Deque.DequeueTail err = %v; want %v", err, nil)
	}

	for v.Len() != 0 {
		v.DequeueTail()
	}

	err, _ = v.DequeueTail()

	if err == nil {
		t.Errorf("Deque.DequeueTail err = %v; want `no tail`", err)
	}
}
