package chapter12

import (
	"strings"
	"testing"
)

func TestSubseqCover(t *testing.T) {
	subseqCoverTests := []struct {
		text     string
		keywords []string
		expected string
	}{
		{
			`My paramount object in this struggle is to save the Union and is not either to save or to destroy slavery`,
			[]string{"Union", "save"},
			"Union and is not either to save",
		},
		{
			`My paramount save object in this save struggle is to save the Union`,
			[]string{"save", "Union"},
			"save the Union",
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
			[]string{"this", "in", "struggle"},
			"this in this this struggle",
		},
		{
			`My paramount save object in this Union save struggle is to save the Union`,
			[]string{"save", "Union"},
			"save the Union",
		},
		{
			`My paramount object in this this in this this struggle is to save the Union`,
			[]string{"Union", "save"},
			"",
		},
	}

	for _, tt := range subseqCoverTests {
		actual := SubseqCover(strings.Split(tt.text, " "), tt.keywords)
		if strings.Join(actual, " ") != tt.expected {
			t.Errorf("SubseqCover(%s, %v): expected %s, actual %s",
				tt.text, tt.keywords, tt.expected, strings.Join(actual, " "))
		}
	}
}
