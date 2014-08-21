package chapter9

import "testing"

func TestKBalancedBinaryTree(t *testing.T) {
	tree := TreeElement{314,
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

	actual := KBalancedBinaryTree(&tree, 3)

	if actual == nil {
		t.Fatalf("KBalancedBinaryTree(%v, %d): unexpected nil value", tree, 3)
	} else if actual.Value != 2 {
		t.Fatalf("KBalancedBinaryTree(%v, %d): expected %d, actual %d", tree, 3, 2, actual.Value)
	}
}
