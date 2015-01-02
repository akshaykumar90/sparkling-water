package common

import "testing"

func AssertIntsAreEqual(t *testing.T, handle string, expected, actual []int) {
	if len(expected) != len(actual) {
		t.Fatalf("%s: expected length %d, actual length %d", handle, len(expected), len(actual))
	}

	for i, v := range expected {
		if v != actual[i] {
			t.Fatalf("%s: at %d: expected %d, actual %d", handle, i, v, actual[i])
		}
	}
}
