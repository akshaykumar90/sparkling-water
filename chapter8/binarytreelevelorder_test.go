package chapter8

import "testing"

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

	expected := [][]int{
		{19},
		{7, 43},
		{3, 11, 23},
	}

	assertEqualLists := func(level int, expected, actual []int) {
		if len(expected) != len(actual) {
			t.Fatalf("BinaryTreeLevelOrder: at level %d: expected length %d, actual length %d", level, len(expected), len(actual))
		}

		for i, v := range expected {
			if v != actual[i] {
				t.Fatalf("BinaryTreeLevelOrder: at level %d: at %d: expected %d, actual %d", level, i, v, actual[i])
			}
		}
	}

	actual := BinaryTreeLevelOrder(&tree)

	if len(expected) != len(actual) {
		t.Fatalf("BinaryTreeLevelOrder: levels: expected length %d, actual length %d", len(expected), len(actual))
	}

	for i, v := range expected {
		assertEqualLists(i+1, v, actual[i])
	}
}

func TestEmptyBinaryTreeLevelOrder(t *testing.T) {
	actual := BinaryTreeLevelOrder(nil)
	if len(actual) != 0 {
		t.Fatalf("BinaryTreeLevelOrder(nil): expected length %d, actual length %d", 0, len(actual))
	}
}
