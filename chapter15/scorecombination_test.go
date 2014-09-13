package chapter15

import "testing"

var scoreCombinationTests = []struct {
	W        []int
	s        int
	expected int
}{
	{[]int{2, 3, 7}, 12, 4},
}

func TestScoreCombination(t *testing.T) {
	for _, tt := range scoreCombinationTests {
		actual := ScoreCombination(tt.W, tt.s)
		if actual != tt.expected {
			t.Errorf("ScoreCombination(%v, %d): expected %d, actual %d",
				tt.W, tt.s, tt.expected, actual)
		}
	}
}
