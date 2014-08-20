package chapter8

import "testing"

func TestSlidingWindow(t *testing.T) {
	traffic := []TrafficElement{{0, 0}, {1, 1}, {2, 3}, {3, 1}, {4, 0}, {5, 2}, {6, 2}, {7, 2}}
	expected := []int{0, 1, 3, 3, 3, 3, 2, 2}
	actual := SlidingWindow(traffic, 3)

	if len(expected) != len(actual) {
		t.Fatalf("SlidingWindow(%v, %d): expected length %d, actual length %d",
			traffic, 3, len(expected), len(actual))
	} else {
		for i, v := range expected {
			if v != actual[i] {
				t.Fatalf("SlidingWindow(%v, %d): at %d: expected %d, actual %d",
					traffic, 3, i+1, v, actual[i])
			}
		}
	}
}
