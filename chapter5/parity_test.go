package chapter5

import "testing"

var parityTests = []struct {
	n        uint64 // input
	expected uint8  // expected result
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 0},
	{4, 1},
	{1<<64 - 1, 0},
}

func TestParity(t *testing.T) {
	for _, tt := range parityTests {
		actual := Parity(tt.n)
		if actual != tt.expected {
			t.Errorf("Parity(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}
