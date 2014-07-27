package chapter8

import (
	"fmt"
	"testing"
)

func TestBinaryTreeLevelOrder(t *testing.T) {
	tree := BST{19,
		&BST{7,
			&BST{Value: 3},
			&BST{Value: 11},
		},
		&BST{43,
			&BST{Value: 23},
			nil,
		},
	}

	BinaryTreeLevelOrder(&tree)
	fmt.Println()
}

func TestEmptyBinaryTreeLevelOrder(t *testing.T) {
	BinaryTreeLevelOrder(nil)
	fmt.Println()
}