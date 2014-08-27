package chapter9

import "testing"

func TestKThNodeBinaryTree(t *testing.T) {
	root := &TreeElementWithSize{3, 6,
		&TreeElementWithSize{2, 2,
			&TreeElementWithSize{1, 1, nil, nil},
			nil,
		},
		&TreeElementWithSize{5, 3,
			&TreeElementWithSize{4, 1, nil, nil},
			&TreeElementWithSize{6, 1, nil, nil},
		},
	}

	kThNodeBinaryTreeTests := []struct {
		k        int
		expected int
	}{
		{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6},
	}

	for _, tt := range kThNodeBinaryTreeTests {
		actual := KThNodeBinaryTree(root, tt.k).Value
		if actual != tt.expected {
			t.Errorf("KThNodeBinaryTree(%d): expected %d, actual %d",
				tt.k, tt.expected, actual)
		}
	}
}
