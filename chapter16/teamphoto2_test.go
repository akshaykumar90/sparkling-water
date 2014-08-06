package chapter16

import (
	"testing"
)

func TestTeamPhoto2(t *testing.T) {
	th := [][]int{
		{3, 0, 2},
		{2, 3, 4},
		{2, 1, 0},
		{1, 5, 4},
		{3, 5, 6},
	}

	actual := TeamPhoto2(th)
	if actual != 3 {
		t.Errorf("TeamPhoto2(%v): expected %d, actual %d",
			th, 3, actual)
	}
}
