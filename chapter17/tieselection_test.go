package chapter17

import "testing"

func TestTiesElection(t *testing.T) {
	V := []int{
		9, 3, 11, 6, 55, 9, 7, 3, 29, 16, 4, 4, 20, 11, 6, 6, 8,
		8, 4, 10, 11, 16, 10, 6, 10, 3, 5, 6, 4, 14, 5, 29, 15, 3,
		18, 7, 7, 20, 4, 9, 3, 11, 38, 6, 3, 13, 12, 5, 10, 3, 3,
	}

	actual := TiesElection(V)
	expected := 16976480564070
	if actual != expected {
		t.Errorf("TiesElection: expected %d, actual %d", expected, actual)
	}
}
