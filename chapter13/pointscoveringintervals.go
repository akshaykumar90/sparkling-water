// problem 13.12

package chapter13

import "sort"

type Time int

type Interval struct {
	start, end Time
}

type TimeType int

const (
	Start TimeType = iota
	End
)

type IntervalTime struct {
	id       int
	time     Time
	timeType TimeType
}

type IntervalTimeSlice []IntervalTime

func (p IntervalTimeSlice) Len() int           { return len(p) }
func (p IntervalTimeSlice) Less(i, j int) bool { return p[i].time < p[j].time }
func (p IntervalTimeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func PointsCoveringIntervals(a []Interval) []Time {
	timeSlice := make(IntervalTimeSlice, len(a)*2)

	result := make([]Time, 0)

	for i, interval := range a {
		timeSlice[2*i] = IntervalTime{i, interval.start, Start}
		timeSlice[2*i+1] = IntervalTime{i, interval.end, End}
	}

	sort.Sort(timeSlice)

	curr := make(map[int]bool)

	for _, intervalTime := range timeSlice {
		switch intervalTime.timeType {
		case Start:
			curr[intervalTime.id] = true
		case End:
			if curr[intervalTime.id] {
				result = append(result, intervalTime.time)
				curr = make(map[int]bool)
			}
		}
	}

	return result
}
