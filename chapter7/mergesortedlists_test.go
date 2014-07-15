package chapter7

import "testing"

func TestMergeSortedLists(t *testing.T) {
	fst := Element{2, &Element{5, &Element{8, nil}}}
	snd := Element{1, &Element{4, nil}}

	expected := &Element{1, &Element{2, &Element{4, &Element{5, &Element{8, nil}}}}}
	actual := MergeSortedLists(&fst, &snd)

	l := 0
	for ; actual != nil && expected != nil; l++ {
		if actual.Value != expected.Value {
			t.Errorf("MergeSortedLists: mismatch, expected %d actual %d", expected.Value, actual.Value)
		}
		actual = actual.Next
		expected = expected.Next
	}

	if actual != nil || expected != nil {
		t.Errorf("MergeSortedLists: premature termination, matching prefix length %d", l)
	}
}
