package chapter9

import "testing"

// Figure 9.1: Example of a binary tree
var ExampleBinaryTree = &TreeElement{314,
	&TreeElement{6,
		&TreeElement{271,
			&TreeElement{28, nil, nil},
			&TreeElement{0, nil, nil},
		},
		&TreeElement{561,
			nil,
			&TreeElement{3,
				&TreeElement{17, nil, nil},
				nil,
			},
		},
	},
	&TreeElement{6,
		&TreeElement{2,
			nil,
			&TreeElement{1,
				&TreeElement{401,
					nil,
					&TreeElement{641, nil, nil},
				},
				&TreeElement{257, nil, nil},
			},
		},
		&TreeElement{271,
			nil,
			&TreeElement{28, nil, nil},
		},
	},
}

func TestConnectLeavesBinaryTree(t *testing.T) {

	expected := []int{28, 0, 17, 641, 257, 28}

	actual := ConnectLeavesBinaryTree(ExampleBinaryTree)

	if len(expected) != len(actual) {
		t.Fatalf("ConnectLeavesBinaryTree(%v): expected length %d, actual length %d",
			ExampleBinaryTree, len(expected), len(actual))
	} else {
		for i, v := range expected {
			if v != actual[i].Value {
				t.Fatalf("ConnectLeavesBinaryTree(%v): at %d: expected %d, actual %d",
					ExampleBinaryTree, i, v, actual[i].Value)
			}
		}
	}
}
