package chapter7

import "testing"

func TestZippingListEvenLength(t *testing.T) {
	list := &Element{1, &Element{2, &Element{3, &Element{4, &Element{5, &Element{6, nil}}}}}}

	expected := &Element{1, &Element{6, &Element{2, &Element{5, &Element{3, &Element{4, nil}}}}}}

	actual := ZippingList(list)

	assertEqualLists(t, "ZippingList", expected, actual)
}

func TestZippingListOddLength(t *testing.T) {
	list := &Element{1, &Element{2, &Element{3, &Element{4, &Element{5, nil}}}}}

	expected := &Element{1, &Element{5, &Element{2, &Element{4, &Element{3, nil}}}}}

	actual := ZippingList(list)

	assertEqualLists(t, "ZippingList", expected, actual)
}

func TestZippingListNil(t *testing.T) {
	assertEqualLists(t, "ZippingList", nil, ZippingList(nil))
}

func TestZippingListSingleElement(t *testing.T) {
	assertEqualLists(t, "ZippingList", &Element{1, nil}, ZippingList(&Element{1, nil}))
}
