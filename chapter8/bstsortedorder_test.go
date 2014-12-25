package chapter8

import "testing"

func TestBSTSortedOrder(t *testing.T) {
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

	expected := []int{3, 7, 11, 19, 23, 43}

	actual := BSTSortedOrder(&tree)

	if len(expected) != len(actual) {
		t.Fatalf("BSTSortedOrder: expected length %d, actual length %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if v != actual[i] {
			t.Fatalf("BSTSortedOrder: at %d: expected %d, actual %d", i, v, actual[i])
		}
	}
}
