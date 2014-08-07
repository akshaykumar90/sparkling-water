package chapter17

import "testing"

func TestZeroOneKnapsack(t *testing.T) {
	items := []item{
		{20, 65},
		{8, 35},
		{60, 245},
		{55, 195},
		{40, 65},
		{70, 150},
		{85, 275},
		{25, 155},
		{30, 120},
		{65, 320},
		{75, 75},
		{10, 40},
		{95, 200},
		{50, 100},
		{40, 220},
		{10, 99},
	}

	actual := ZeroOneKnapsack(items, 130)
	expected := 695
	if actual != expected {
		t.Errorf("ZeroOneKnapsack(%v, %d): expected %d, actual %d",
			items, 130, expected, actual)
	}
}
