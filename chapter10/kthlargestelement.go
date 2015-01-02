// Problem 10.5

package chapter10

import (
	"container/heap"
	"github.com/akshaykumar90/sparkling-water/common"
)

func KThLargestElement(stream []int, k int) []int {
	result := make([]int, 0, len(stream))

	h := make(common.IntHeap, 0)

	heap.Init(&h)

	for i := 0; i < k && i < len(stream); i++ {
		heap.Push(&h, stream[i])
		result = append(result, h[0])
	}

	for i := k; i < len(stream); i++ {
		if stream[i] > h[0] {
			heap.Pop(&h)
			heap.Push(&h, stream[i])
		}
		result = append(result, h[0])
	}

	return result
}
