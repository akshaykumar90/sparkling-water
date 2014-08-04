package chapter15

import "testing"

func TestTreeDiameter(t *testing.T) {
	n1 := WeightedTreeNode{Value: 1}
	n2 := WeightedTreeNode{Value: 2}
	n3 := WeightedTreeNode{Value: 3}

	n4 := WeightedTreeNode{2, []*WeightedTreeNode{&n1, &n2, &n3}}
	n5 := WeightedTreeNode{Value: 4}

	n6 := WeightedTreeNode{4, []*WeightedTreeNode{&n4, &n5}}
	n7 := WeightedTreeNode{Value: 6}

	n8 := WeightedTreeNode{1, []*WeightedTreeNode{&n6, &n7}}
	n9 := WeightedTreeNode{Value: 2}

	n10 := WeightedTreeNode{3, []*WeightedTreeNode{&n8, &n9}}

	n11 := WeightedTreeNode{Value: 14}

	n12 := WeightedTreeNode{Value: 6}

	n13 := WeightedTreeNode{4, []*WeightedTreeNode{&n12}}
	n14 := WeightedTreeNode{Value: 3}

	n15 := WeightedTreeNode{7, []*WeightedTreeNode{&n13, &n14}}

	root := WeightedTreeNode{0, []*WeightedTreeNode{&n15, &n11, &n10}}

	actual := TreeDiameter(&root)
	if actual != 31 {
		t.Errorf("TreeDiameter(%v): expected %d, actual %d",
			root, 31, actual)
	}

}
