package chapter6

import "testing"

var regularExpressionTests = []struct {
	r, s     string
	expected bool
}{
	{".", "a", true},
	{"a", "a", true},
	{"a", "b", false},
	{"a.9", "aW9", true},
	{"a.9", "aW19", false},
	{"^a.9", "aW9", true},
	{"^a.9", "baW19", false},
	{".*", "a", true},
	{".*", "", true},
	{"c*", "c", true},
	{"aa*", "c", false},
	{"ca*", "c", true},
	{".*", "asdsdsa", true},
	{"9$", "xxxxW19", true},
	{".*a", "ba", true},
	{".*a", "ba", true},
	{"^a.*9$", "axxxxW19", true},
	{"^a.*W19$", "axxxxW19", true},
	{".*a.*W19", "axxxxW19123", true},
	{".*b.*W19", "axxxxW19123", false},
	{"n.*", "n", true},
	{"a*n.*", "an", true},
	{"a*n.*", "ana", true},
	{"a*n.*W19", "anaxxxxW19123", true},
	{".*a*n.*W19", "asdaaadnanaxxxxW19123", true},
	{".*.*.a*n.*W19", "asdaaadnanaxxxxW19123", true},
}

func TestRegularExpression(t *testing.T) {
	for _, tt := range regularExpressionTests {
		actual := RegularExpression(tt.r, tt.s)
		if actual != tt.expected {
			t.Errorf("RegularExpression(%q, %q): expected %t, actual %t",
				tt.r, tt.s, tt.expected, actual)
		}
	}
}
