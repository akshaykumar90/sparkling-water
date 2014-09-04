package chapter12

import (
	"strings"
	"testing"
)

func TestSmallestSubarrayCoveringSet(t *testing.T) {
	smallestSubarrayCoveringSetTests := []struct {
		text     string
		keywords []string
		expected string
	}{
		{
			`My paramount save object in this save struggle is to save the Union`,
			[]string{"Union", "save"},
			"save the Union",
		},
		{
			`My paramount save object in this save Union struggle is to save the Union`,
			[]string{"Union", "save"},
			"save Union",
		},
		{
			`My paramount save object in this Union save struggle is to save the Union`,
			[]string{"Union", "save"},
			"Union save",
		},
		{
			`My paramount object in this struggle is to save the Union`,
			[]string{"object", "struggle", "save"},
			"object in this struggle is to save",
		},
		{
			`My paramount object in this this in this this struggle is to save the Union`,
			[]string{"in", "this", "struggle"},
			"in this this struggle",
		},
		{
			`My paramount object in this this in this this struggle is to save the Union`,
			[]string{"hello", "world"},
			"",
		},
	}

	for _, tt := range smallestSubarrayCoveringSetTests {
		actual := SmallestSubarrayCoveringSet(strings.Split(tt.text, " "), tt.keywords)
		if strings.Join(actual, " ") != tt.expected {
			t.Errorf("SmallestSubarrayCoveringSet(%s, %v): expected %s, actual %s",
				tt.text, tt.keywords, tt.expected, strings.Join(actual, " "))
		}
	}
}
