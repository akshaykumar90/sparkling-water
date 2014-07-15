package chapter7

import "testing"

func TestCheckCycle(t *testing.T) {
	cycle_start := Element{3, nil}
	head := Element{1, &Element{2, &cycle_start}}
	cycle_start.Next = &Element{4, &Element{5, &Element{6, &Element{7, &cycle_start}}}}

	actual := CheckCycle(&head)

	if actual != &cycle_start {
		t.Errorf("TestCheckCycle: expected %s actual %s", cycle_start, actual)
	}
}
