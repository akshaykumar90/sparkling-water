package chapter9

import "testing"

func TestInorderTraversalWithParent(t *testing.T) {
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

	expected := []int{3, 7, 11, 19, 23, 43}

	actual := InorderTraversalWithParent(&n19)

	if len(expected) != len(actual) {
		t.Fatalf("InorderTraversalWithParent: expected length %d, actual length %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if v != actual[i] {
			t.Fatalf("InorderTraversalWithParent: at %d: expected %d, actual %d", i, v, actual[i])
		}
	}
}
