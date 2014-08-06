package chapter13

import "testing"

var teamPhoto1Tests = []struct {
	a        []int
	b        []int
	expected bool
}{
	{[]int{1, 5, 4}, []int{2, 3, 4}, false},
	{[]int{2, 3, 4}, []int{1, 5, 4}, false},
	{[]int{0, 3, 2}, []int{1, 5, 4}, true},
	{[]int{1, 5, 4}, []int{0, 3, 2}, false},
	{[]int{0, 3, 2}, []int{2, 3, 4}, true},
}

func TestTeamPhoto1(t *testing.T) {
	for _, tt := range teamPhoto1Tests {
		actual := TeamPhoto1(tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("TeamPhoto1(%v, %v): expected %t, actual %t",
				tt.a, tt.b, tt.expected, actual)
		}
	}
}
