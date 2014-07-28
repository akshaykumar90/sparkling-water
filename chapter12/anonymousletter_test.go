package chapter12

import "testing"

var anonymousLetterTests = []struct {
	L        string
	M        string
	expected bool
}{
	{"hello", "why hello there!", true},
	{"how are you?", "very well", false},
	{"this is a letter", "", false},
	{"epi", "epic", true},
}

func TestAnonymousLetter(t *testing.T) {
	for _, tt := range anonymousLetterTests {
		actual := AnonymousLetter(tt.L, tt.M)
		if actual != tt.expected {
			t.Errorf("AnonymousLetter(%s, %s): expected %t, actual %t",
				tt.L, tt.M, tt.expected, actual)
		}
	}
}
