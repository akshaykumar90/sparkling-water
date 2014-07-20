package chapter9

import "testing"

func TestLowestCommonAncestorHash(t *testing.T) {
	n19 := TreeNodeWithParent{Value: 19}
	n7 := TreeNodeWithParent{Value: 7}
	n11 := TreeNodeWithParent{Value: 11}
	n3 := TreeNodeWithParent{Value: 3}
	n43 := TreeNodeWithParent{Value: 43}
	n23 := TreeNodeWithParent{Value: 23}

	n19.Left, n19.Right = &n7, &n43
	n7.Left, n7.Right = &n3, &n11
	n43.Left = &n23

	n3.Parent = &n7
	n11.Parent = &n7

	n23.Parent = &n43

	n7.Parent = &n19
	n43.Parent = &n19

	var lowestCommonAncestorHashTests = []struct {
		fst      *TreeNodeWithParent
		snd      *TreeNodeWithParent
		expected *TreeNodeWithParent
	}{
		{&n3, &n11, &n7},
		{&n3, &n7, &n7},
		{&n11, &n23, &n19},
		{&n19, &n3, &n19},
	}

	for _, tt := range lowestCommonAncestorHashTests {
		actual := LowestCommonAncestorHash(tt.fst, tt.snd)
		if actual != tt.expected {
			t.Errorf("LowestCommonAncestorHash(%d, %d): expected %d, actual %v",
				tt.fst.Value, tt.snd.Value, tt.expected.Value, actual)
		}
	}
}
