package chapter6

import "testing"

var bigNumberMultiplicationTests = []struct {
	fst      string
	snd      string
	expected int
}{
	{"224", "36", 224 * 36},
	{"-456", "1234", -456 * 1234},
	{"1355", "-2411", 1355 * -2411},
	{"-421", "-532", -421 * -532},
	{"0", "0", 0},
	{"0", "-1", 0},
	{"-1", "-1", 1},
	{"231", "0", 0},
	{"-1", "531", -531},
}

func TestBigNumberMultiplication(t *testing.T) {
	for _, tt := range bigNumberMultiplicationTests {
		actual := BigNumberMultiplication(tt.fst, tt.snd)
		if actual != tt.expected {
			t.Errorf("BigNumberMultiplication(%s, %s): expected %d, actual %d",
				tt.fst, tt.snd, tt.expected, actual)
		}
	}
}
