package chapter6

import "testing"

var reverseWordsTests = []struct {
	str      string
	expected string
}{
	{"Bob likes Alice", "Alice likes Bob"},
	{"This is a Test", "Test a is This"},
	{"Double Word", "Word Double"},
	{"Single", "Single"},
}

func TestReverseWords(t *testing.T) {
	for _, tt := range reverseWordsTests {
		actual := ReverseWords(tt.str)
		if actual != tt.expected {
			t.Errorf("ReverseWords(%s): expected %s, actual %s",
				tt.str, tt.expected, actual)
		}
	}
}
