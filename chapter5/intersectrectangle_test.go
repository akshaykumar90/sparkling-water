package chapter5

import "testing"

var validIntersectRectangleTests = []struct {
	r1       Rectangle
	r2       Rectangle
	expected Rectangle
}{
	{Rectangle{0, 0, 2, 1}, Rectangle{1, 0, 2, 1}, Rectangle{1, 0, 1, 1}},
	{Rectangle{0, 0, 1, 2}, Rectangle{0, 1, 1, 2}, Rectangle{0, 1, 1, 1}},
	{Rectangle{0, 0, 2, 2}, Rectangle{1, 1, 2, 2}, Rectangle{1, 1, 1, 1}},
	{Rectangle{1, 1, 2, 2}, Rectangle{0, 0, 2, 2}, Rectangle{1, 1, 1, 1}},
	{Rectangle{0, 0, 2, 2}, Rectangle{1, -1, 2, 2}, Rectangle{1, 0, 1, 1}},
	{Rectangle{0, 0, 1, 1}, Rectangle{0, 0, 1, 1}, Rectangle{0, 0, 1, 1}},
	{Rectangle{0, 0, 2, 2}, Rectangle{1, 1, 1, 1}, Rectangle{1, 1, 1, 1}},
}

var invalidIntersectRectangleTests = []struct {
	r1 Rectangle
	r2 Rectangle
}{
	{Rectangle{0, 0, 1, 1}, Rectangle{1, 0, 1, 1}},
	{Rectangle{0, 0, 1, 1}, Rectangle{0, 1, 1, 1}},
	{Rectangle{0, 0, 1, 1}, Rectangle{2, 0, 1, 1}},
	{Rectangle{0, 0, 1, 1}, Rectangle{0, 2, 1, 1}},
	{Rectangle{0, 0, 1, 1}, Rectangle{2, 2, 2, 2}},
	{Rectangle{0, 0, 2, 2}, Rectangle{1, 3, 1, 1}},
	{Rectangle{0, 0, 2, 2}, Rectangle{1, 4, 1, 1}},
}

func TestIntersectRectangleValid(t *testing.T) {
	for _, tt := range validIntersectRectangleTests {
		actual := IntersectRectangle(tt.r1, tt.r2)
		if *actual != tt.expected {
			t.Errorf("IntersectRectangle(%v, %v): expected %v, actual %v",
				tt.r1, tt.r2, tt.expected, *actual)
		}
	}
}

func TestIntersectRectangleInvalid(t *testing.T) {
	for _, tt := range invalidIntersectRectangleTests {
		actual := IntersectRectangle(tt.r1, tt.r2)
		if actual != nil {
			t.Errorf("IntersectRectangle(%v, %v): expected nil, actual %v",
				tt.r1, tt.r2, *actual)
		}
	}
}
