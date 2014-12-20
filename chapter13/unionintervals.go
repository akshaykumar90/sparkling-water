// Problem 13.11

package chapter13

import (
	"fmt"
	"sort"
)

type OpenClosedInterval struct {
	start, end Endpoint
}

type Endpoint struct {
	time     Time
	isClosed bool
}

func (interval OpenClosedInterval) String() string {
	left, right := '(', ')'

	if interval.start.isClosed {
		left = '['
	}

	if interval.end.isClosed {
		right = ']'
	}

	return fmt.Sprintf("%c%d,%d%c", left, interval.start.time, interval.end.time, right)
}

type ByStartTime []OpenClosedInterval

func (s ByStartTime) Len() int { return len(s) }

func (s ByStartTime) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s ByStartTime) Less(i, j int) bool {
	if s[i].start.time != s[j].start.time {
		return s[i].start.time < s[j].start.time
	}
	return s[i].start.isClosed
}

func UnionIntervals(intervals []OpenClosedInterval) []OpenClosedInterval {
	if len(intervals) == 0 {
		return []OpenClosedInterval{}
	}

	union := make([]OpenClosedInterval, 0)

	sort.Sort(ByStartTime(intervals))

	cand := intervals[0]

	for _, x := range intervals[1:] {
		switch {
		case x.start.time > cand.end.time:
			union = append(union, cand)
			cand = x
		case x.start.time == cand.end.time:
			if x.start.isClosed || cand.end.isClosed {
				cand.end = x.end
			} else {
				union = append(union, cand)
				cand = x
			}
		default: // x.start.time < cand.end.time
			if x.end.time > cand.end.time || (x.end.time == cand.end.time && x.end.isClosed) {
				cand.end = x.end
			}
		}
	}

	union = append(union, cand)

	return union
}
