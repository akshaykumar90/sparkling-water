//problem 10.6
package chapter10

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func ApproximateSort(in []int, k int) []int {
	out := make([]int, 0)

	h := make(IntHeap, k+1)

	copy(h, in[:k+1])

	heap.Init(&h)

	for _, v := range in[k+1:] {
		out = append(out, heap.Pop(&h).(int))
		heap.Push(&h, v)
	}

	for len(h) > 0 {
		out = append(out, heap.Pop(&h).(int))
	}

	return out
}
