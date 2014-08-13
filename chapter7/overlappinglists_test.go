package chapter7

import "testing"

func TestOverlappingLists00(t *testing.T) {
	common_list := Element{4, &Element{5, &Element{6, &Element{7, nil}}}}
	fst := Element{1, &Element{2, &common_list}}
	snd := Element{3, &common_list}

	actual := OverlappingLists(&fst, &snd)

	if actual != &common_list {
		t.Errorf("TestOverlappingLists00: expected %v actual %v", common_list, actual)
	}
}

func TestOverlappingLists01(t *testing.T) {
	fst := Element{8, &Element{9, nil}}

	cycle_start := Element{3, nil}
	snd := Element{1, &Element{2, &cycle_start}}
	cycle_start.Next = &Element{4, &Element{5, &Element{6, &Element{7, &cycle_start}}}}

	actual := OverlappingLists(&fst, &snd)

	if actual != nil {
		t.Errorf("TestOverlappingLists01: expected nil actual %v", actual)
	}
}

func TestOverlappingLists10(t *testing.T) {
	cycle_start := Element{3, nil}
	fst := Element{1, &Element{2, &cycle_start}}
	cycle_start.Next = &Element{4, &Element{5, &Element{6, &Element{7, &cycle_start}}}}

	snd := Element{8, &Element{9, nil}}

	actual := OverlappingLists(&fst, &snd)

	if actual != nil {
		t.Errorf("TestOverlappingLists10: expected nil actual %v", actual)
	}
}

func TestOverlappingLists11(t *testing.T) {
	cycle_start_fst := Element{3, nil}
	cycle_start_snd := Element{5, nil}

	cycle_start_fst.Next = &Element{4, &cycle_start_snd}
	cycle_start_snd.Next = &Element{6, &Element{7, &cycle_start_fst}}

	fst := Element{1, &Element{2, &cycle_start_fst}}
	snd := Element{8, &Element{9, &cycle_start_snd}}

	actual := OverlappingLists(&fst, &snd)

	if !(actual == &cycle_start_fst || actual == &cycle_start_snd) {
		t.Errorf("TestOverlappingLists11: expected one of cycle start point: actual %v", actual)
	}
}
