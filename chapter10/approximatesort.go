// Problem 10.6

package chapter10

import (
	"container/heap"
	"github.com/akshaykumar90/sparkling-water/common"
)

func ApproximateSort(in []int, k int) []int {
	out := make([]int, 0)

	h := make(common.IntHeap, k)

	copy(h, in[:k])

	heap.Init(&h)

	for _, v := range in[k:] {
		heap.Push(&h, v)
		out = append(out, heap.Pop(&h).(int))
	}

	for len(h) > 0 {
		out = append(out, heap.Pop(&h).(int))
	}

	return out
}
