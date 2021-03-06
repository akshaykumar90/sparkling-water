// Problem 10.8

package chapter10

import (
	"container/heap"
	"github.com/akshaykumar90/sparkling-water/common"
)

func OnlineMedianHelper(a []int) (common.IntHeap, common.IntHeap, []float64) {
	minHeap, maxHeap := make(common.IntHeap, 0), make(common.IntHeap, 0)

	medians := []float64{float64(a[0])}

	maxHeap = append(maxHeap, -a[0])

	heap.Init(&maxHeap)
	heap.Init(&minHeap)

	for i := 1; i < len(a); i++ {
		if a[i] <= -maxHeap[0] {
			heap.Push(&maxHeap, -a[i])
			if i&1 != 0 {
				heap.Push(&minHeap, -heap.Pop(&maxHeap).(int))
			}
		} else {
			heap.Push(&minHeap, a[i])
			if i&1 == 0 {
				heap.Push(&maxHeap, -heap.Pop(&minHeap).(int))
			}
		}

		if (i+1)&1 == 0 {
			medians = append(medians, float64(minHeap[0]-maxHeap[0])/2.0)
		} else {
			medians = append(medians, float64(-maxHeap[0]))
		}
	}

	return maxHeap, minHeap, medians
}

func OnlineMedian(s []int) []float64 {
	_, _, medians := OnlineMedianHelper(s)
	return medians
}
