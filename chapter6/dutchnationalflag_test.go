package chapter6

import "testing"

func TestDutchNationalFlag(t *testing.T) {
	var dutchNationalFlagTests = []struct {
		a []int
		p int
	}{
		{[]int{1}, 0},
		{[]int{1, 2, 3}, 1},
		{[]int{5, 3, 9, 3, 8, 1}, 3},
	}

	for _, tt := range dutchNationalFlagTests {
		res := make([]int, len(tt.a))
		copy(res, tt.a)
		DutchNationalFlag(res, tt.p)

		i := 0
		for ; res[i] < tt.p; i++ {
		}

		for ; res[i] == tt.p; i++ {
		}

		for ; i < len(res); i++ {
			if res[i] <= tt.p {
				t.Fatalf("DutchNationalFlag(%v, %d) -> %v :at %d: %d <= %d",
					tt.a, tt.p, res, i, res[i], tt.p)
			}
		}
	}
}
