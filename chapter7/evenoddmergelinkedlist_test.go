package chapter7

import "testing"

func assertEqualLists(t *testing.T, name string, expected, actual *Element) {
	l := 0
	for ; actual != nil && expected != nil; l++ {
		if actual.Value != expected.Value {
			t.Errorf("%s: mismatch, expected %d actual %d", name, expected.Value, actual.Value)
		}
		actual = actual.Next
		expected = expected.Next
	}

	if actual != nil || expected != nil {
		t.Errorf("%s: premature termination, matching prefix length %d", name, l)
	}
}

func TestEvenOddMergeLinkedList(t *testing.T) {
	list := Element{1, &Element{2, &Element{3, &Element{4, &Element{5, nil}}}}}

	expected := &Element{1, &Element{3, &Element{5, &Element{2, &Element{4, nil}}}}}

	actual := EvenOddMergeLinkedList(&list)

	assertEqualLists(t, "EvenOddMergeLinkedList", expected, actual)
}

func TestEvenOddMergeLinkedListNil(t *testing.T) {
	assertEqualLists(t, "EvenOddMergeLinkedList", nil, EvenOddMergeLinkedList(nil))
}

func TestEvenOddMergeLinkedListSingleElement(t *testing.T) {
	assertEqualLists(t, "EvenOddMergeLinkedList", &Element{1, nil}, EvenOddMergeLinkedList(&Element{1, nil}))
}

func TestEvenOddMergeLinkedListTwoElements(t *testing.T) {
	assertEqualLists(t, "EvenOddMergeLinkedList", &Element{1, &Element{2, nil}},
		EvenOddMergeLinkedList(&Element{1, &Element{2, nil}}))
}

func TestEvenOddMergeLinkedListThreeElements(t *testing.T) {
	assertEqualLists(t, "EvenOddMergeLinkedList", &Element{1, &Element{3, &Element{2, nil}}},
		EvenOddMergeLinkedList(&Element{1, &Element{2, &Element{3, nil}}}))
}

func TestEvenOddMergeLinkedListFourElement(t *testing.T) {
	assertEqualLists(t, "EvenOddMergeLinkedList", &Element{1, &Element{3, &Element{2, &Element{4, nil}}}},
		EvenOddMergeLinkedList(&Element{1, &Element{2, &Element{3, &Element{4, nil}}}}))
}
