package chapter9

import "testing"

var shortestUniquePrefixTests = []struct {
	s        string
	d        []string
	expected string
}{
	{"cat", []string{"dog", "be"}, "c"},
	{"cat", []string{"dog", "be", "cut"}, "ca"},
	{"cat", []string{"dog", "be", "cut", "car"}, "cat"},
}

func TestShortestUniquePrefix(t *testing.T) {
	for _, tt := range shortestUniquePrefixTests {
		actual := ShortestUniquePrefix(tt.s, tt.d)
		if actual != tt.expected {
			t.Errorf("ShortestUniquePrefix(%q, %v): expected %q, actual %q",
				tt.s, tt.d, tt.expected, actual)
		}
	}
}
