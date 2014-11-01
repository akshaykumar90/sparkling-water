package chapter9

import "testing"

// Figure 9.2: Symmetric and asymmetric binary trees.
var symmetricBinaryTreeTests = []struct {
	root     *TreeElement
	expected bool
}{
	{
		&TreeElement{314,
			&TreeElement{6, nil, &TreeElement{2, nil, &TreeElement{Value: 3}}},
			&TreeElement{6, &TreeElement{2, &TreeElement{Value: 3}, nil}, nil},
		}, true,
	},
	{
		&TreeElement{314,
			&TreeElement{6, nil, &TreeElement{561, nil, &TreeElement{Value: 3}}},
			&TreeElement{6, &TreeElement{2, &TreeElement{Value: 1}, nil}, nil},
		}, false,
	},
	{
		&TreeElement{314,
			&TreeElement{6, nil, &TreeElement{561, nil, &TreeElement{Value: 3}}},
			&TreeElement{6, &TreeElement{Value: 561}, nil},
		}, false,
	},
}

func TestSymmetricBinaryTree(t *testing.T) {
	for _, tt := range symmetricBinaryTreeTests {
		actual := SymmetricBinaryTree(tt.root)
		if actual != tt.expected {
			t.Errorf("SymmetricBinaryTree(%v): expected %t, actual %t",
				tt.root, tt.expected, actual)
		}
	}
}
