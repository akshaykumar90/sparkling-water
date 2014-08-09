package chapter5

import "testing"

var spreadsheetEncodingTests = []struct {
	col      string
	expected int
}{
	{"A", 1},
	{"B", 2},
	{"Y", 25},
	{"Z", 26},
	{"AA", 27},
	{"AB", 28},
	{"ZY", 701},
	{"ZZ", 702},
	{"BZ", 78},
}

func TestSpreadsheetEncoding(t *testing.T) {
	for _, tt := range spreadsheetEncodingTests {
		actual := SpreadsheetEncoding(tt.col)
		if actual != tt.expected {
			t.Errorf("SpreadsheetEncoding(%s): expected %d, actual %d",
				tt.col, tt.expected, actual)
		}
	}
}
