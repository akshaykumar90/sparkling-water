package chapter13

import (
	"fmt"
	"testing"
)

func TestUnionIntervals(t *testing.T) {
	// Figure 13.3 - A set of intervals and their union.
	intervals := []OpenClosedInterval{
		{Endpoint{0, false}, Endpoint{3, false}},
		{Endpoint{1, true}, Endpoint{1, true}},
		{Endpoint{2, true}, Endpoint{4, true}},
		{Endpoint{3, true}, Endpoint{4, false}},
		{Endpoint{5, true}, Endpoint{7, false}},
		{Endpoint{7, true}, Endpoint{8, false}},
		{Endpoint{8, true}, Endpoint{11, false}},
		{Endpoint{9, false}, Endpoint{11, true}},
		{Endpoint{12, true}, Endpoint{14, true}},
		{Endpoint{12, false}, Endpoint{15, true}},
		{Endpoint{13, false}, Endpoint{13, false}},
		{Endpoint{15, false}, Endpoint{17, false}},
	}

	fmt.Println("TestUnionIntervals:", UnionIntervals(intervals))
}
