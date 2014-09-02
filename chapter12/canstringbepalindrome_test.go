package chapter12

import "testing"

var canStringBePalindromeTests = []struct {
	s        string
	expected bool
}{
	{"level", true},
	{"rotator", true},
	{"edified", true},
	{"explore", false},
	{"", true},
	{"space", false},
}

func TestCanStringBePalindrome(t *testing.T) {
	for _, tt := range canStringBePalindromeTests {
		actual := CanStringBePalindrome(tt.s)
		if actual != tt.expected {
			t.Errorf("CanStringBePalindrome(%s): expected %t, actual %t",
				tt.s, tt.expected, actual)
		}
	}
}
