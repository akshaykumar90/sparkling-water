// problem 10.5

package chapter10

import (
	"container/heap"
	"fmt"
)

func KThLargestElement(stream []int, k int) {
	h := make(IntHeap, 0)
	heap.Init(&h)
	for i := 0; i < k && i < len(stream); i++ {
		heap.Push(&h, stream[i])
		fmt.Printf("kth largest after %d cycles: %d\n", i+1, h[0])
	}
	for i := k; i < len(stream); i++ {
		if stream[i] > h[0] {
			heap.Pop(&h)
			heap.Push(&h, stream[i])
		}
		fmt.Printf("kth largest after %d cycles: %d\n", i+1, h[0])
	}
}
