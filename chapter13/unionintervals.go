// Problem 13.11

package chapter13

type OpenClosedInterval struct {
	start, end Endpoint
}

type Endpoint struct {
	time     Time
	isClosed bool
}

func UnionIntervals(I []OpenClosedInterval) []OpenClosedInterval {
	return []OpenClosedInterval{}
}
