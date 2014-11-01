package chapter9

import "testing"

func TestKBalancedBinaryTree(t *testing.T) {
	actual := KBalancedBinaryTree(ExampleBinaryTree, 3)

	if actual == nil {
		t.Fatalf("KBalancedBinaryTree(%v, %d): unexpected nil value", ExampleBinaryTree, 3)
	} else if actual.Value != 2 {
		t.Fatalf("KBalancedBinaryTree(%v, %d): expected %d, actual %d", ExampleBinaryTree, 3, 2, actual.Value)
	}
}
