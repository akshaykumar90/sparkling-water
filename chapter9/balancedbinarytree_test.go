package chapter9

import "testing"

func TestBalancedBinaryTree(t *testing.T) {
	actual := BalancedBinaryTree(ExampleBinaryTree)
	if actual != false {
		t.Fatalf("BalancedBinaryTree: expected %t, actual %t", false, actual)
	}
}
