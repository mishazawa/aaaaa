package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()

	testval1 := "first"
	testval2 := "second"
	testval3 := "third"

	s.Push(testval1)

	err, val := s.Peek()

	if err != nil {
		t.Errorf("Stack.Peek = %v; want %v", err, nil)
	}

	if val != testval1 {
		t.Errorf("Stack.Peek = %v; want %v", val, testval1)
	}

	s.Push(testval2)
	s.Push(testval3)

	err, val = s.Peek()

	if err != nil {
		t.Errorf("Stack.Peek = %v; want %v", err, nil)
	}

	if val != testval3 {
		t.Errorf("Stack.Peek = %v; want %v", val, testval3)
	}

	err, val = s.Pop()

	if err != nil {
		t.Errorf("Stack.Pop = %v; want %v", err, nil)
	}

	if val != testval3 {
		t.Errorf("Stack.Pop = %v; want %v", val, testval3)
	}

	s.Pop()
	s.Pop()

	err, _ = s.Pop()

	if err == nil {
		t.Errorf("Stack.Pop = %v; want 'no tail'", err)
	}
}
