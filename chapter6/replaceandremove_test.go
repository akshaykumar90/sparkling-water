package chapter6

import "testing"

var replaceAndRemoveTests = []struct {
	str      string
	expected string
}{
	{"daccbbabc", "dddccddc"},
	{"bbbbb", ""},
	{"aaaa", "dddddddd"},
	{"bbaa", "dddd"},
	{"abab", "dddd"},
	{"abcddcba", "ddcddcdd"},
	{"cadccabbbdda", "cdddccdddddd"},
}

func TestReplaceAndRemove(t *testing.T) {
	for _, tt := range replaceAndRemoveTests {
		actual := ReplaceAndRemove(tt.str)
		if actual != tt.expected {
			t.Errorf("ReplaceAndRemove(%s): expected %s, actual %s",
				tt.str, tt.expected, actual)
		}
	}
}
