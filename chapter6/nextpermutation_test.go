package chapter6

import "testing"

var validNextPermutationTests = []struct {
	p        []int
	expected []int
}{
	{
		[]int{3, 4},
		[]int{4, 3},
	},
	{
		[]int{1, 0, 3, 2},
		[]int{1, 2, 0, 3},
	},
	{
		[]int{4, 2, 9},
		[]int{4, 9, 2},
	},
	{
		[]int{4, 5, 3, 2},
		[]int{5, 2, 3, 4},
	},
}

var invalidNextPermutationTests = [][]int{
	{1},
	{2, 1},
	{3, 2, 1},
}

func TestValidNextPermutation(t *testing.T) {
	for _, tt := range validNextPermutationTests {
		cc := make([]int, len(tt.p))
		copy(cc, tt.p)
		actual := NextPermutation(cc)

		for i, v := range tt.expected {
			if v != actual[i] {
				t.Errorf("NextPermutation(%v): at %d: expected %d, actual %d",
					tt.p, i+1, v, actual[i])
			}
		}
	}
}

func TestInvalidNextPermutation(t *testing.T) {
	for _, tt := range invalidNextPermutationTests {
		cc := make([]int, len(tt))
		copy(cc, tt)
		actual := NextPermutation(cc)

		if actual != nil {
			t.Errorf("NextPermutation(%v): expected nil, actual %v", tt, actual)
		}
	}
}
