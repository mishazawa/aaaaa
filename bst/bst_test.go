package bst

import "testing"

type TestWrapInt int

func NewTestWrapInt(v int) TestWrapInt {
	return TestWrapInt(v)
}

func (a TestWrapInt) Gte(b interface{}) bool {
	aint := int(a)
	bint := toInt(b)
	return aint >= bint
}

func toInt(v interface{}) int {
	return int(v.(TestWrapInt))
}

/*
  When inserting values in following order 3, 5, 1, 6, 2, 4, 7, 8, 9
  tree should look like this

       3
      / \
     1   5
    / \  |\
 nil   2 | \
         |  6
         4   \
              7 - 8 - 9
*/

func TestBstAdd(t *testing.T) {
	testval1 := 1
	testval2 := 2
	testval3 := 3
	testval4 := 4
	testval5 := 5
	testval6 := 6
	testval7 := 7
	testval8 := 8
	testval9 := 9

	root := NewRoot(NewTestWrapInt(testval3))

	if toInt(root.Value) != testval3 {
		t.Errorf("root.Value = %v; want = %v", root.Value, testval3)
	}

	root.Add(NewTestWrapInt(testval5))

	if root.Right == nil {
		t.Errorf("root.Right = %v; want = %v", root.Right, NewTestWrapInt(testval5))
	}

	if root.Left != nil {
		t.Errorf("root.Left = %v; want = %v", root.Left, nil)
	}

	if toInt(root.Right.Value) != testval5 {
		t.Errorf("root.Value = %v; want = %v", root.Right.Value, testval5)
	}

	root.Add(NewTestWrapInt(testval1))
	root.Add(NewTestWrapInt(testval6))
	root.Add(NewTestWrapInt(testval2))
	root.Add(NewTestWrapInt(testval4))
	root.Add(NewTestWrapInt(testval7))
	root.Add(NewTestWrapInt(testval8))
	root.Add(NewTestWrapInt(testval9))

	// valid leafs

	leaf1 := root.Left.Value
	leaf6 := root.Right.Right.Value
	leaf2 := root.Left.Right.Value
	leaf4 := root.Right.Left.Value
	leaf7 := root.Right.Right.Right.Value
	leaf8 := root.Right.Right.Right.Right.Value
	leaf9 := root.Right.Right.Right.Right.Right.Value

	if toInt(leaf4) != testval4 {
		t.Errorf("leaf 4 = %v; want = %v", leaf4, testval4)
	}

	if toInt(leaf1) != testval1 {
		t.Errorf("leaf 1 = %v; want = %v", leaf1, testval1)
	}

	if toInt(leaf2) != testval2 {
		t.Errorf("leaf 2 = %v; want = %v", leaf2, testval2)
	}

	if toInt(leaf6) != testval6 {
		t.Errorf("leaf 6 = %v; want = %v", leaf6, testval6)
	}

	if toInt(leaf7) != testval7 {
		t.Errorf("leaf 7 = %v; want = %v", leaf7, testval7)
	}

	if toInt(leaf8) != testval8 {
		t.Errorf("leaf 8 = %v; want = %v", leaf8, testval8)
	}

	if toInt(leaf9) != testval9 {
		t.Errorf("leaf 9 = %v; want = %v", leaf9, testval9)
	}

	if root.Left.Left != nil {
		t.Errorf("leaf root.Left.Left = %v; want = %v", root.Left.Left, nil)
	}

	root.Add(NewTestWrapInt(0))

	if toInt(root.Left.Left.Value) != 0 {
		t.Errorf("leaf root.Left.Left = %v; want = %v", root.Left.Left.Value, 0)
	}

}