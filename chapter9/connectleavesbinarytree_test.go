package chapter9

import "testing"

func TestConnectLeavesBinaryTree(t *testing.T) {
	tree := &TreeElement{314,
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

	expected := []int{28, 0, 17, 641, 257, 28}

	actual := ConnectLeavesBinaryTree(tree)

	if len(expected) != len(actual) {
		t.Fatalf("ConnectLeavesBinaryTree(%v): expected length %d, actual length %d",
			tree, len(expected), len(actual))
	} else {
		for i, v := range expected {
			if v != actual[i].Value {
				t.Fatalf("ConnectLeavesBinaryTree(%v): at %d: expected %d, actual %d",
					tree, i, v, actual[i].Value)
			}
		}
	}
}
