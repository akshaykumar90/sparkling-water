package chapter9

import "testing"

var pre = []string{"H", "B", "F", "null", "null", "E", "A", "null", "null", "null", "C", "null", "D", "null", "G", "I", "null", "null", "null"}

func TestReconstructPreorderWithNull(t *testing.T) {
	assertEqualBinaryTrees(ReconstructThisBinaryTree, ReconstructPreorderWithNull(pre))
}

func TestReconstructPreorderWithNull2(t *testing.T) {
	assertEqualBinaryTrees(ReconstructThisBinaryTree, ReconstructPreorderWithNull2(pre))
}
