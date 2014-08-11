package chapter6

import "testing"

var robotBatteryTests = []struct {
	hs       []int
	expected int
}{
	{[]int{1, 2, 3, 4}, 3},
	{[]int{1, 1, 1}, 0},
	{[]int{5, 3, 1, -1}, 0},
	{[]int{8, 3, 1, -1, -10, 2, 7, 17, 32, 14, 9}, 42},
}

func TestRobotBattery(t *testing.T) {
	for _, tt := range robotBatteryTests {
		actual := RobotBattery(tt.hs)
		if actual != tt.expected {
			t.Errorf("RobotBattery(%v): expected %d, actual %d",
				tt.hs, tt.expected, actual)
		}
	}
}
