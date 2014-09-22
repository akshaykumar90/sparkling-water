package chapter15

import "testing"

var computeBinomialCoefficientsTests = []struct {
	n        int
	k        int
	expected int
}{
	{8, 0, 1},
	{10, 10, 1},
	{1, 0, 1},
	{1, 1, 1},
	{2, 1, 2},
	{26, 14, 9657700},
	{31, 10, 44352165},
	{47, 6, 10737573},
}

func TestComputeBinomialCoefficients(t *testing.T) {
	for _, tt := range computeBinomialCoefficientsTests {
		actual := ComputeBinomialCoefficients(tt.n, tt.k)
		if actual != tt.expected {
			t.Errorf("ComputeBinomialCoefficients(%d, %d): expected %d, actual %d",
				tt.n, tt.k, tt.expected, actual)
		}
	}
}
