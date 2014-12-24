package chapter9

import "testing"

func TestExteriorBinaryTree(t *testing.T) {
	expected := []int{314, 6, 271, 28, 0, 17, 641, 257, 28, 271, 6}

	actual := ExteriorBinaryTree(ExampleBinaryTree)

	if len(expected) != len(actual) {
		t.Fatalf("ExteriorBinaryTree: expected length %d, actual length %d", len(expected), len(actual))
	} else {
		for i, v := range expected {
			if v != actual[i] {
				t.Fatalf("ExteriorBinaryTree: at %d: expected %d, actual %d", i, v, actual[i])
			}
		}
	}
}
