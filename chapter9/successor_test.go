package chapter9

import "testing"

func TestSuccessor(t *testing.T) {
	n19 := TreeNodeWithParent{Value: 19}
	n7 := TreeNodeWithParent{Value: 7}
	n11 := TreeNodeWithParent{Value: 11}
	n3 := TreeNodeWithParent{Value: 3}
	n43 := TreeNodeWithParent{Value: 43}
	n23 := TreeNodeWithParent{Value: 23}

	n19.Left, n19.Right = &n7, &n43
	n7.Left, n7.Right = &n3, &n11
	n43.Left = &n23

	n3.Parent = &n7
	n11.Parent = &n7

	n23.Parent = &n43

	n7.Parent = &n19
	n43.Parent = &n19

	var successorTests = []struct {
		node     *TreeNodeWithParent
		expected *TreeNodeWithParent
	}{
		{&n3, &n7},
		{&n7, &n11},
		{&n11, &n19},
		{&n19, &n23},
		{&n23, &n43},
		{&n43, nil},
	}

	for _, tt := range successorTests {
		actual := Successor(tt.node)
		if actual != tt.expected {
			t.Errorf("Successor(%v): expected %v, actual %v", tt.node, tt.expected, actual)
		}
	}
}
