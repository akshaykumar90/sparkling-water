//problem 10.8
package chapter10

import (
	"container/heap"
	"fmt"
)

func OnlineMedian(s []int) {
	minHeap, maxHeap := make(IntHeap, 0), make(IntHeap, 0)

	seen := make([]int, 0)

	median := float64(s[0])

	maxHeap = append(maxHeap, -s[0])

	heap.Init(&maxHeap)
	heap.Init(&minHeap)

	seen = append(seen, s[0])
	fmt.Println(seen)
	fmt.Println(median)

	for _, v := range s[1:] {
		if v < -maxHeap[0] {
			heap.Push(&maxHeap, -v)
		} else {
			heap.Push(&minHeap, v)
		}

		switch len(minHeap) - len(maxHeap) {
		case -2:
			heap.Push(&minHeap, -heap.Pop(&maxHeap).(int))
			median = float64(minHeap[0]-maxHeap[0]) / 2.0
		case +2:
			heap.Push(&maxHeap, -heap.Pop(&minHeap).(int))
			median = float64(minHeap[0]-maxHeap[0]) / 2.0
		case -1:
			median = float64(-maxHeap[0])
		case +1:
			median = float64(minHeap[0])
		case 0:
			median = float64(minHeap[0]-maxHeap[0]) / 2.0
		}

		seen = append(seen, v)
		fmt.Println(seen)
		fmt.Println(median)

	}

}
