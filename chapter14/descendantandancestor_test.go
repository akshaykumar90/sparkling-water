package chapter14

import "testing"

var descendantAndAncestorTests = []struct {
	r, s, m  *TreeNode
	expected bool
}{
	{
		tree,
		tree.Right.Left.Right,
		tree.Right.Left,
		true,
	},
	{
		tree.Right,
		tree.Right.Right.Right,
		tree.Right.Left,
		false,
	},
	{
		tree.Right.Left,
		tree.Right.Right,
		tree.Right.Left.Right,
		false,
	},
	{
		tree.Left.Left,
		tree.Left.Left.Left,
		tree.Left,
		false,
	},
}

func TestDescendantAndAncestor(t *testing.T) {
	for _, tt := range descendantAndAncestorTests {
		actual := DescendantAndAncestor(tt.r, tt.s, tt.m)
		if actual != tt.expected {
			t.Errorf("DescendantAndAncestor(%d, %d, %d): expected %t, actual %t",
				tt.r.Value, tt.s.Value, tt.m.Value, tt.expected, actual)
		}
	}
}
