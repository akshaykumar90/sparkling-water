package chapter5

import "testing"

var convertBaseTests = []struct {
	repr     string
	b1       int
	b2       int
	expected string
}{
	{"39374091", 10, 16, "258CD0B"},
	{"39374091", 10, 8, "226146413"},
	{"39374091", 10, 2, "10010110001100110100001011"},
	{"16D34", 16, 8, "266464"},
	{"16D34", 16, 10, "93492"},
	{"0", 10, 2, "0"},
	{"0", 10, 16, "0"},
	{"30031", 8, 8, "30031"},
}

func TestConvertBase(t *testing.T) {
	for _, tt := range convertBaseTests {
		actual := ConvertBase(tt.repr, tt.b1, tt.b2)
		if actual != tt.expected {
			t.Errorf("ConvertBase(%s, %d, %d): expected %s, actual %s",
				tt.repr, tt.b1, tt.b2, tt.expected, actual)
		}
	}
}
