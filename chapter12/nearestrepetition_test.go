package chapter12

import (
	"math"
	"testing"
)

var nearestRepetitionTests = []struct {
	S        []string
	expected int
}{
	{[]string{"foo", "bar", "widget", "foo", "widget", "widget"}, 1},
	{[]string{"foo", "bar", "widget", "foo", "xyz", "widget", "bar"}, 3},
	{[]string{"foo", "bar", "widget"}, math.MaxInt32},
	{[]string{"foo", "foo", "foo"}, 1},
	{[]string{}, math.MaxInt32},
}

func TestNearestRepetition(t *testing.T) {
	for _, tt := range nearestRepetitionTests {
		actual := NearestRepetition(tt.S)
		if actual != tt.expected {
			t.Errorf("NearestRepetition(%v): expected %d, actual %d",
				tt.S, tt.expected, actual)
		}
	}
}
