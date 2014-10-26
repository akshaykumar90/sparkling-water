package chapter15

import "testing"

func TestPickingUpCoins(t *testing.T) {
	actual, expected := PickingUpCoins([]int{25, 5, 10, 5, 10, 5, 10, 25, 1, 25, 1, 25, 1, 25, 5, 10}), 140
	if actual != expected {
		t.Errorf("PickingUpCoins: expected %d, actual %d", expected, actual)
	}
}
