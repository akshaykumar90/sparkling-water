package chapter14

import "testing"

func TestViewFromAbove(t *testing.T) {
	// Figure 14.7: Instance of the view from above problem, with patterns used to denote colors.
	seq := []LineSegment{
		{0, 4, 1, 1},
		{5, 7, 2, 1},
		{9, 18, 3, 1},
		{2, 7, 4, 2},
		{8, 9, 1, 2},
		{12, 15, 3, 2},
		{1, 3, 5, 3},
		{6, 10, 1, 3},
		{11, 13, 2, 3},
		{14, 15, 4, 3},
		{16, 17, 2, 3},
		{4, 5, 2, 4},
	}

	expected := []LineSegment{
		{0, 1, 1, 1},
		{1, 3, 5, 3},
		{3, 4, 4, 2},
		{4, 5, 2, 4},
		{5, 6, 4, 2},
		{6, 10, 1, 3},
		{10, 11, 3, 1},
		{11, 13, 2, 3},
		{13, 14, 3, 2},
		{14, 15, 4, 3},
		{15, 16, 3, 1},
		{16, 17, 2, 3},
		{17, 18, 3, 1},
	}

	actual := ViewFromAbove(seq)

	if len(expected) != len(actual) {
		t.Fatalf("ViewFromAbove: expected length %d, actual length %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if v != actual[i] {
			t.Fatalf("ViewFromAbove: at %d: expected %d, actual %d", i, v, actual[i])
		}
	}
}
