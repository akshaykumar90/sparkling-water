package chapter7

import "testing"

func TestOverlappingListsNoCycle(t *testing.T) {
	common_list := Element{4, &Element{5, &Element{6, &Element{7, nil}}}}
	fst := Element{1, &Element{2, &common_list}}
	snd := Element{3, &common_list}

	actual := OverlappingListsNoCycle(&fst, &snd)

	if actual != &common_list {
		t.Errorf("OverlappingListsNoCycle: expected %s actual %s", common_list, actual)
	}
}
