package chapter9

import "testing"

var lowestCommonAncestorNoParentTests = []struct {
	a, b     *TreeElement
	expected *TreeElement
}{
	{
		ExampleBinaryTree.Right.Left.Right.Left.Right,
		ExampleBinaryTree.Right.Left.Right.Right,
		ExampleBinaryTree.Right.Left.Right,
	},
	{
		ExampleBinaryTree.Right,
		ExampleBinaryTree.Left,
		ExampleBinaryTree,
	},
	{
		ExampleBinaryTree.Right,
		ExampleBinaryTree,
		ExampleBinaryTree,
	},
}

func TestLowestCommonAncestorNoParent(t *testing.T) {
	for _, tt := range lowestCommonAncestorNoParentTests {
		actual := LowestCommonAncestorNoParent(ExampleBinaryTree, tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("LowestCommonAncestorNoParent(%d, %d): expected %d, actual %d",
				tt.a.Value, tt.b.Value, tt.expected.Value, actual.Value)
		}
	}
}
