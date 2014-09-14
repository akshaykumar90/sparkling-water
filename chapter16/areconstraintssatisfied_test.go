package chapter16

import "testing"

var areConstraintsSatisfiedTests = []struct {
	eq       []Pair
	ieq      []Pair
	expected bool
}{
	{
		[]Pair{{1, 2}, {2, 3}, {3, 4}},
		[]Pair{{1, 4}},
		false,
	},
	{
		[]Pair{{1, 2}, {3, 4}},
		[]Pair{{2, 3}},
		true,
	},
	{
		[]Pair{{1, 4}, {4, 2}, {3, 5}},
		[]Pair{{3, 4}, {4, 5}, {1, 3}},
		true,
	},
}

func TestAreConstraintsSatisfied(t *testing.T) {
	for _, tt := range areConstraintsSatisfiedTests {
		actual := AreConstraintsSatisfied(tt.eq, tt.ieq)
		if actual != tt.expected {
			t.Errorf("AreConstraintsSatisfied(%v, %v): expected %t, actual %t",
				tt.eq, tt.ieq, tt.expected, actual)
		}
	}
}
