package queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()

	testval1 := "first"
	testval2 := "second"
	testval3 := "third"

	q.Enqueue(testval1)

	err, val := q.Peek()

	if err != nil {
		t.Errorf("Queue.Peek = %v; want %v", err, nil)
	}

	if val != testval1 {
		t.Errorf("Queue.Peek = %v; want %v", val, testval1)
	}

	q.Enqueue(testval2)
	q.Enqueue(testval3)

	err, val = q.Peek()

	if err != nil {
		t.Errorf("Queue.Peek = %v; want %v", err, nil)
	}

	if val != testval1 {
		t.Errorf("Queue.Peek = %v; want %v", val, testval1)
	}

	err, val = q.Dequeue()

	if err != nil {
		t.Errorf("Queue.Dequeue = %v; want %v", err, nil)
	}

	if val != testval1 {
		t.Errorf("Queue.Dequeue = %v; want %v", val, testval1)
	}

	q.Dequeue() // val 2
	q.Dequeue() // val 3

	err, _ = q.Dequeue()

	if err == nil {
		t.Errorf("Queue.Dequeue = %v; want 'no head'", err)
	}
}
