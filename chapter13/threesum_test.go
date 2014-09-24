package chapter13

import "testing"

func TestThreeSum(t *testing.T) {
	a := []int{62, 47, 60, 63, 25, 67, 30, 24, 73, 58}

	threeSumTests := []struct {
		t        int
		expected bool
	}{
		{110, true},
		{125, false},
		{200, true},
		{220, false},
		{72, true},
		{81, false},
	}

	for _, tt := range threeSumTests {
		actual := ThreeSum(a, tt.t)
		if actual != tt.expected {
			t.Errorf("ThreeSum(%d): expected %t, actual %t",
				tt.t, tt.expected, actual)
		}
	}
}
