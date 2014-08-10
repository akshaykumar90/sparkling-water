package chapter5

import "testing"

func TestPrimeSieve(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	actual := PrimeSieve(100)

	if len(expected) != len(actual) {
		t.Fatalf("PrimeSieve(100): expected length %d, actual length %d",
			len(expected), len(actual))
	} else {
		for i, v := range expected {
			if v != actual[i] {
				t.Fatalf("PrimeSieve(100): at %d: expected %d, actual %d",
					i+1, v, actual[i])
			}
		}
	}
}
