package chapter15

import "testing"

var levenshteinDistanceTests = []struct {
	a        string
	b        string
	expected int
}{
	{"Carthorse", "Orchestra", 8},
	{"thou shalt not", "you should not", 5},
}

func TestLevenshteinDistance(t *testing.T) {
	for _, tt := range levenshteinDistanceTests {
		actual := LevenshteinDistance(tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("LevenshteinDistance(%s, %s): expected %d, actual %d",
				tt.a, tt.b, tt.expected, actual)
		}
	}
}
