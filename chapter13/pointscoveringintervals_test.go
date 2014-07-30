package chapter13

import "testing"

var pointsCoveringIntervalsTests = []struct {
	a        []Interval
	expected int
}{
	{[]Interval{{0, 1}}, 1},
	{[]Interval{{0, 1}, {2, 3}}, 2},
	{[]Interval{{0, 2}, {1, 5}, {4, 6}}, 2},
	{[]Interval{{0, 6}, {1, 5}, {2, 4}}, 1},
}

func TestPointsCoveringIntervals(t *testing.T) {
	for _, tt := range pointsCoveringIntervalsTests {
		actual := len(PointsCoveringIntervals(tt.a))
		if actual != tt.expected {
			t.Errorf("PointsCoveringIntervals(%v): expected %d, actual %d",
				tt.a, tt.expected, actual)
		}
	}
}
