package chapter9

import (
	"fmt"
	"testing"
)

func TestReconstructPreorderWithNull(t *testing.T) {
	pre := []string{"H", "B", "F", "null", "null", "E", "A", "null", "null", "null", "C", "null", "D", "null", "G", "I", "null", "null", "null"}

	root := ReconstructPreorderWithNull(pre)

	fmt.Println("TestReconstructPreorderWithNull")
	inOrderTraversal(root)
	fmt.Println()
}
