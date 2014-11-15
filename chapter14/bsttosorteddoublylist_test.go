package chapter14

import "testing"

func TestBSTToSortedDoublyList(t *testing.T) {
	// Figure 14.4: BST to sorted doubly linked list.
	_7 := &TreeNode{Value: 7}
	_3 := &TreeNode{Value: 3}
	_11 := &TreeNode{Value: 11}
	_2 := &TreeNode{Value: 2}
	_5 := &TreeNode{Value: 5}

	_7.Left, _7.Right = _3, _11
	_3.Left, _3.Right = _2, _5

	root := _7

	expected := []*TreeNode{_2, _3, _5, _7, _11}

	actual, _ := BSTToSortedDoublyList(root)

	for i, v := range expected {
		if actual != v {
			t.Fatalf("BSTToSortedDoublyList: at %d: expected %d, actual %v", i+1, v.Value, actual.Value)
		}
		actual = actual.Right
	}
}
