// Problem 14.19

package chapter14

import (
	"github.com/petar/GoLLRB/llrb"
	"sort"
)

type LineSegment struct {
	l, r, c, h int
}

func (p *LineSegment) Less(than llrb.Item) bool {
	return p.h < than.(*LineSegment).h
}

type endpoint struct {
	x      int
	seg    *LineSegment
	isLeft bool
}

type endpointList []endpoint

func (lst endpointList) Len() int           { return len(lst) }
func (lst endpointList) Less(i, j int) bool { return lst[i].x < lst[j].x }
func (lst endpointList) Swap(i, j int)      { lst[i], lst[j] = lst[j], lst[i] }

func ViewFromAbove(seq []LineSegment) []LineSegment {
	allEndpoints := make(endpointList, 0, len(seq)*2)

	for i := range seq {
		allEndpoints = append(allEndpoints, endpoint{seq[i].l, &seq[i], true}, endpoint{seq[i].r, &seq[i], false})
	}

	sort.Sort(allEndpoints)

	active := llrb.New()

	result := make([]LineSegment, 0)

	var curr, prevMax *LineSegment

	for _, p := range allEndpoints {
		if p.isLeft {
			active.ReplaceOrInsert(p.seg)
		} else {
			active.Delete(p.seg)
		}

		var max *LineSegment
		if active.Len() > 0 {
			max = active.Max().(*LineSegment)
		}

		if curr == nil {
			curr = &LineSegment{l: p.x, c: max.c, h: max.h}
		} else if max != prevMax {
			curr.r = p.x
			if curr.r != curr.l {
				result = append(result, *curr)
			}
			if max == nil {
				curr = nil
			} else {
				curr = &LineSegment{l: p.x, c: max.c, h: max.h}
			}
		}

		prevMax = max
	}

	return result
}
