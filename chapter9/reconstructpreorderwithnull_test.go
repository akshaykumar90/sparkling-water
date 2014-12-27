package chapter9

import "testing"

var pre = []string{"H", "B", "F", "null", "null", "E", "A", "null", "null", "null", "C", "null", "D", "null", "G", "I", "null", "null", "null"}

func TestReconstructPreorderWithNull(t *testing.T) {
	if !assertEqualBinaryTrees(ReconstructThisBinaryTree, ReconstructPreorderWithNull(pre)) {
		t.Fatalf("ReconstructPreorderWithNull: assertEqualBinaryTrees failed!")
	}

}

func TestReconstructPreorderWithNull2(t *testing.T) {
	if !assertEqualBinaryTrees(ReconstructThisBinaryTree, ReconstructPreorderWithNull2(pre)) {
		t.Fatalf("ReconstructPreorderWithNull2: assertEqualBinaryTrees failed!")
	}
}
