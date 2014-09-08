package chapter14

import "testing"

// Figure 14.1 - An example BST.
var tree TreeNode = TreeNode{19,
	&TreeNode{7,
		&TreeNode{3,
			&TreeNode{Value: 2},
			&TreeNode{Value: 5},
		},
		&TreeNode{11,
			nil,
			&TreeNode{17,
				&TreeNode{Value: 13},
				nil,
			},
		},
	},
	&TreeNode{43,
		&TreeNode{23,
			nil,
			&TreeNode{37,
				&TreeNode{29,
					nil,
					&TreeNode{Value: 31},
				},
				&TreeNode{Value: 41},
			},
		},
		&TreeNode{47,
			nil,
			&TreeNode{Value: 53},
		},
	},
}

var validSearchBSTFirstLargerKTests = []struct {
	k        int
	expected int
}{
	{23, 29},
	{11, 13},
	{19, 23},
	{2, 3},
	{43, 47},
	{41, 43},
	{17, 19},
}

var invalidSearchBSTFirstLargerKTests = []int{12, 25, 20, 53}

func TestSearchBSTFirstLargerKValid(t *testing.T) {
	for _, tt := range validSearchBSTFirstLargerKTests {
		actual := SearchBSTFirstLargerK(&tree, tt.k)
		if actual == nil {
			t.Errorf("SearchBSTFirstLargerK(%d): expected %d, actual nil",
				tt.k, tt.expected)
		} else if actual.Value != tt.expected {
			t.Errorf("SearchBSTFirstLargerK(%d): expected %d, actual %d",
				tt.k, tt.expected, actual.Value)
		}
	}
}

func TestSearchBSTFirstLargerKInvalid(t *testing.T) {
	for _, tt := range invalidSearchBSTFirstLargerKTests {
		actual := SearchBSTFirstLargerK(&tree, tt)
		if actual != nil {
			t.Errorf("SearchBSTFirstLargerK(%d): expected nil, actual %v",
				tt, actual)
		}
	}
}
