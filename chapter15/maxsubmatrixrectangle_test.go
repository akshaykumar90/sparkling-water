package chapter15

import "testing"

var maxSubmatrixRectangleTests = []struct {
	mat      [][]bool
	expected int
}{
	{
		[][]bool{
			{true, true, false},
			{true, true, true},
		},
		4,
	},
	{
		[][]bool{
			{true, true, false, false},
			{true, true, true, true},
			{true, true, true, true},
		},
		8,
	},
}

func TestMaxSubmatrixRectangle(t *testing.T) {
	for i, tt := range maxSubmatrixRectangleTests {
		actual := MaxSubmatrixRectangle(tt.mat)
		if actual != tt.expected {
			t.Errorf("MaxSubmatrixRectangle [%d]: expected %d, actual %d", i+1, tt.expected, actual)
		}
	}
}

func TestMaxSubmatrixRectangleImproved(t *testing.T) {
	for i, tt := range maxSubmatrixRectangleTests {
		actual := MaxSubmatrixRectangleImproved(tt.mat)
		if actual != tt.expected {
			t.Errorf("MaxSubmatrixRectangleImproved [%d]: expected %d, actual %d", i+1, tt.expected, actual)
		}
	}
}

var maxSubmatrixSquareTests = []struct {
	mat      [][]bool
	expected int
}{
	{
		[][]bool{
			{true, true, false},
			{true, true, true},
		},
		4,
	},
	{
		[][]bool{
			{true, true, false, false},
			{true, true, true, true},
			{true, true, true, true},
		},
		4,
	},
	{
		[][]bool{
			{true, true, false, false},
			{true, false, true, true},
			{true, true, true, false},
		},
		1,
	},
}

func TestMaxSubmatrixSquare(t *testing.T) {
	for i, tt := range maxSubmatrixSquareTests {
		actual := MaxSubmatrixSquare(tt.mat)
		if actual != tt.expected {
			t.Errorf("MaxSubmatrixSquare [%d]: expected %d, actual %d", i+1, tt.expected, actual)
		}
	}
}
