package chapter15

import "testing"

var scorePermutationTests = []struct {
	W        []int
	s        int
	expected int
}{
	{[]int{2, 3, 7}, 12, 18},
}

func TestScorePermutation(t *testing.T) {
	for _, tt := range scorePermutationTests {
		actual := ScorePermutation(tt.W, tt.s)
		if actual != tt.expected {
			t.Errorf("ScorePermutation(%v, %d): expected %d, actual %d",
				tt.W, tt.s, tt.expected, actual)
		}
	}
}
