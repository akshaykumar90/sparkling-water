package chapter17

import "testing"

func TestMinimumSubarrayDifference(t *testing.T) {
	arr := []int{65, 35, 245, 195, 65, 150, 275, 155, 120, 320, 75, 40, 200, 100, 220, 99}

	actual := MinimumSubarrayDifference(arr)

	expected := 1

	if actual != expected {
		t.Errorf("MinimumSubarrayDifference: expected %d, actual %d", expected, actual)
	}
}
