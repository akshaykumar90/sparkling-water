// problem 14.15

package chapter14

import (
	"github.com/petar/GoLLRB/llrb"
	"math"
)

type ArrData struct {
	idx int
	val int
}

func (p ArrData) Less(than llrb.Item) bool {
	if p.val != than.(ArrData).val {
		return p.val < than.(ArrData).val
	} else {
		return p.idx < than.(ArrData).idx
	}
}

func MinimumDistance3SortedArrays(arrs [][]int) int {
	offset := make([]int, len(arrs))
	curr := llrb.New()
	min := math.MaxInt32

	for i := 0; i < len(arrs); i++ {
		curr.ReplaceOrInsert(ArrData{i, arrs[i][0]})
	}

	for {
		if min > curr.Max().(ArrData).val-curr.Min().(ArrData).val {
			min = curr.Max().(ArrData).val - curr.Min().(ArrData).val
		}
		sm := curr.DeleteMin().(ArrData).idx
		offset[sm]++
		if offset[sm] >= len(arrs[sm]) {
			return min
		} else {
			curr.ReplaceOrInsert(ArrData{sm, arrs[sm][offset[sm]]})
		}
	}
}
