package chapter5

import "testing"

var validStringToIntTests = []struct {
	s        string
	expected int
}{
	{"0", 0},
	{"1", 1},
	{"123", 123},
	{"-8", -8},
	{"03213", 3213},
	{"-00033", -33},
	{"-0", 0},
}

var invalidStringToIntTests = []string{
	"", "-", "--123", "123-4", "456ab12", "wqff",
}

func TestStringToIntValid(t *testing.T) {
	for _, tt := range validStringToIntTests {
		actual, _ := StringToInt(tt.s)
		if actual != tt.expected {
			t.Errorf("StringToInt(%s): expected %d, actual %d",
				tt.s, tt.expected, actual)
		}
	}
}

func TestStringToIntInvalid(t *testing.T) {
	for _, tt := range invalidStringToIntTests {
		actual, err := StringToInt(tt)
		if err == nil {
			t.Errorf("StringToInt(%s): expected error, actual %d",
				tt, actual)
		}
	}
}

var intToStringTests = []struct {
	x        int
	expected string
}{
	{0, "0"},
	{1, "1"},
	{123, "123"},
	{8143, "8143"},
	{-8, "-8"},
	{-3213, "-3213"},
}

func TestIntToString(t *testing.T) {
	for _, tt := range intToStringTests {
		actual := IntToString(tt.x)
		if actual != tt.expected {
			t.Errorf("IntToString(%d): expected %s, actual %s",
				tt.x, tt.expected, actual)
		}
	}
}
