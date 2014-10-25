package chapter15

import "testing"

var numberWaysTests = []struct {
	n, m int
}{
	{5, 5},
	{10, 20},
	{1, 1},
	{5, 1},
	{1, 12},
}

func TestNumberWays(t *testing.T) {
	for _, tt := range numberWaysTests {
		actual, expected := NumberWays(tt.n, tt.m), ComputeBinomialCoefficients(tt.n+tt.m-2, tt.n-1)
		if actual != expected {
			t.Errorf("NumberWays(%d, %d): expected %d, actual %d",
				tt.n, tt.m, expected, actual)
		}
	}
}

var numberWaysObstaclesTests = []struct {
	n, m      int
	obstacles [][]bool
	expected  int
}{
	{3, 3, [][]bool{
		{false, false, true},
		{false, false, false},
		{false, true, false},
	}, 2},
	{3, 3, [][]bool{
		{false, false, true},
		{false, false, true},
		{false, true, false},
	}, 0},
	{3, 3, [][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}, 6},
	{2, 2, [][]bool{
		{true, false},
		{false, false},
	}, 0},
}

func TestNumberWaysObstacles(t *testing.T) {
	for _, tt := range numberWaysObstaclesTests {
		actual := NumberWaysObstacles(tt.n, tt.m, tt.obstacles)
		if actual != tt.expected {
			t.Errorf("NumberWays(%d, %d, %v): expected %d, actual %d",
				tt.n, tt.m, tt.obstacles, tt.expected, actual)
		}
	}
}
