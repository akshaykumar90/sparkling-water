package chapter5

import "testing"

func TestReverseBits(t *testing.T) {
	var reverseBitsTests = []uint64{4, 2, 9, 0, 5, 0xFFFFFFFF, 0xABABFFABCDCDCDCD}

	for _, tt := range reverseBitsTests {
		actual := ReverseBits(ReverseBits(tt))
		if actual != tt {
			t.Errorf("ReverseBits(ReverseBits((%d)): expected %d, actual %d",
				tt, tt, actual)
		}
	}
}
