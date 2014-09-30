package chapter6

import "testing"

var encodingTests = []struct {
	in       string
	expected string
}{
	{"aaaabcccaa", "4a1b3c2a"},
}

func TestEncoding(t *testing.T) {
	for _, tt := range encodingTests {
		actual := Encoding(tt.in)
		if actual != tt.expected {
			t.Errorf("Encoding(%q): expected %q, actual %q",
				tt.in, tt.expected, actual)
		}
	}
}

var decodingTests = []struct {
	in       string
	expected string
}{
	{"3e4f2e", "eeeffffee"},
	{"10a4f2e", "aaaaaaaaaaffffee"},
}

func TestDecoding(t *testing.T) {
	for _, tt := range decodingTests {
		actual := Decoding(tt.in)
		if actual != tt.expected {
			t.Errorf("Decoding(%q): expected %q, actual %q",
				tt.in, tt.expected, actual)
		}
	}
}
