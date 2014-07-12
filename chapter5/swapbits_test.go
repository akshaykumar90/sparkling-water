package chapter5

import "testing"

var swapbitsTests = []struct {
	n        uint64
	i        uint
	j        uint
	expected uint64
}{
	{0, 0, 1, 0},
	{2, 0, 1, 1},
	{47, 1, 4, 61},
	{28, 0, 2, 25},
}

func TestSwapBits(t *testing.T) {
	for _, tt := range swapbitsTests {
		actual := SwapBits(tt.n, tt.i, tt.j)
		if actual != tt.expected {
			t.Errorf("SwapBits(%d, %d, %d): expected %d, actual %d",
				tt.n, tt.i, tt.j, tt.expected, actual)
		}
	}
}
