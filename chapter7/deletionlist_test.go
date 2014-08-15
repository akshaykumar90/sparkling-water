package chapter7

import "testing"

func TestDeletionList(t *testing.T) {
	listSuffix := Element{4, &Element{5, nil}}
	elementToDelete := Element{3, &listSuffix}
	listPrefix := Element{1, &Element{2, &elementToDelete}}

	expected := &Element{1, &Element{2, &Element{4, &Element{5, nil}}}}

	DeletionList(&elementToDelete)

	assertEqualLists(t, "DeletionList", expected, &listPrefix)
}
