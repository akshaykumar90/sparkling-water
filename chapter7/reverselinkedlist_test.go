package chapter7

import "testing"

func TestReverseLinkedList(t *testing.T) {
	list := Element{1, &Element{2, &Element{3, &Element{4, &Element{5, nil}}}}}

	expected := &Element{5, &Element{4, &Element{3, &Element{2, &Element{1, nil}}}}}

	actual := ReverseLinkedList(&list)

	assertEqualLists(t, "ReverseLinkedList", expected, actual)
}

func TestReverseLinkedListNil(t *testing.T) {
	assertEqualLists(t, "ReverseLinkedList", nil, ReverseLinkedList(nil))
}

func TestReverseLinkedListSingleElement(t *testing.T) {
	assertEqualLists(t, "ReverseLinkedList", &Element{1, nil}, ReverseLinkedList(&Element{1, nil}))
}

func TestReverseLinkedListTwoElements(t *testing.T) {
	assertEqualLists(t, "ReverseLinkedList", &Element{1, &Element{2, nil}},
		ReverseLinkedList(&Element{2, &Element{1, nil}}))
}
