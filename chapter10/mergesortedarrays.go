// Problem 10.1

package chapter10

import "container/heap"

type Pair struct {
	idx, val int
}

type PairHeap []Pair

func (h PairHeap) Len() int {
	return len(h)
}

func (h PairHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h PairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	var x Pair
	x, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return x
}

func MergeSortedArrays(arrs [][]int) []int {
	h := new(PairHeap)

	heap.Init(h)

	for i, file := range arrs {
		heap.Push(h, Pair{i, file[0]})
	}

	var result []int
	offset := make([]int, len(arrs))

	for len(*h) > 0 {
		top := heap.Pop(h).(Pair)
		result = append(result, top.val)
		offset[top.idx]++
		if offset[top.idx] < len(arrs[top.idx]) {
			heap.Push(h, Pair{top.idx, arrs[top.idx][offset[top.idx]]})
		}
	}

	return result

}
