package chapter6

import "testing"

var permutationArrayTests = []struct {
	P, A     []int
	expected []int
}{
	{
		[]int{3, 1, 0, 2},
		[]int{2, 1, 3, 0},
		[]int{0, 1, 2, 3},
	},
}

func TestPermutationArray(t *testing.T) {
	for _, tt := range permutationArrayTests {
		ca := make([]int, len(tt.A))
		copy(ca, tt.A)

		actual := PermutationArray(tt.P, ca)

		for i, v := range tt.expected {
			if v != actual[i] {
				t.Errorf("PermutationArray(%v, %v): at %d: expected %d, actual %d",
					tt.P, tt.A, i+1, v, actual[i])
			}
		}
	}
}
