// problem 13.10

package chapter13

import "sort"

func RenderingCalendar(a []Interval) int {
	timeSlice := make(IntervalTimeSlice, len(a)*2)

	for i, interval := range a {
		timeSlice[2*i] = IntervalTime{i, interval.start, Start}
		timeSlice[2*i+1] = IntervalTime{i, interval.end, End}
	}

	sort.Sort(timeSlice)

	curr, max := 0, 0

	for _, it := range timeSlice {
		if it.timeType == Start {
			curr++
		} else {
			curr--
		}
		if max < curr {
			max = curr
		}
	}

	return max
}
