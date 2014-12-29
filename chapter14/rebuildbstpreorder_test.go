package chapter14

import "testing"

func assertEqualBinaryTrees(expected, actual *TreeNode) bool {
	if expected == nil && actual == nil {
		return true
	} else if expected != nil && actual != nil {
		return expected.Value == actual.Value &&
			assertEqualBinaryTrees(expected.Left, actual.Left) &&
			assertEqualBinaryTrees(expected.Right, actual.Right)
	} else {
		return false
	}
}

func TestRebuildBSTPreorder(t *testing.T) {
	actual := RebuildBSTPreorder([]int{19, 7, 3, 2, 5, 11, 17, 13, 43, 23, 37, 29, 31, 41, 47, 53})

	if !assertEqualBinaryTrees(&tree, actual) {
		t.Fatalf("RebuildBSTPreorder: assertEqualBinaryTrees failed!")
	}
}

func TestRebuildBSTPreorderBetter(t *testing.T) {
	actual := RebuildBSTPreorderBetter([]int{19, 7, 3, 2, 5, 11, 17, 13, 43, 23, 37, 29, 31, 41, 47, 53})

	if !assertEqualBinaryTrees(&tree, actual) {
		t.Fatalf("RebuildBSTPreorderBetter: assertEqualBinaryTrees failed!")
	}
}
