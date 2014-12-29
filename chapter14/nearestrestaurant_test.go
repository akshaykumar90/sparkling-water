package chapter14

import "testing"

var nearestRestaurantTests = []struct {
	lo, hi   int
	expected []int
}{
	{10, 20, []int{11, 13, 17, 19}},
	{20, 40, []int{23, 29, 31, 37}},
	{47, 53, []int{47, 53}},
	{11, 11, []int{11}},
	{60, 70, []int{}},
}

func TestNearestRestaurant(t *testing.T) {
	for _, tt := range nearestRestaurantTests {
		actual := NearestRestaurant(tree, tt.lo, tt.hi)

		if len(tt.expected) != len(actual) {
			t.Fatalf("NearestRestaurant(%d,%d): expected length %d, actual length %d",
				tt.lo, tt.hi, len(tt.expected), len(actual))
		}

		for i, v := range tt.expected {
			if v != actual[i] {
				t.Fatalf("NearestRestaurant(%d,%d): at %d: expected %d, actual %d",
					tt.lo, tt.hi, i, v, actual[i])
			}
		}
	}
}
