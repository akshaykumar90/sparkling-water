package chapter13

import "testing"

var renderingCalendarTests = []struct {
	a        []Interval
	expected int
}{
	{[]Interval{{0, 1}}, 1},
	{[]Interval{{0, 1}, {2, 3}}, 1},
	{[]Interval{{0, 2}, {1, 5}, {4, 6}}, 2},
	{[]Interval{{0, 6}, {1, 5}, {2, 4}}, 3},
}

func TestRenderingCalendar(t *testing.T) {
	for _, tt := range renderingCalendarTests {
		actual := RenderingCalendar(tt.a)
		if actual != tt.expected {
			t.Errorf("RenderingCalendar(%v): expected %d, actual %d",
				tt.a, tt.expected, actual)
		}
	}
}
