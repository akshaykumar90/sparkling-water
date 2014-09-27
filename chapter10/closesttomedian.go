// Problem 10.7

package chapter10

import "container/heap"

func ClosestToMedian(a []int, k int) []int {
	leftHeap, rightHeap, medians := OnlineMedianHelper(a)

	median := medians[len(medians)-1]

	closest := make([]int, k)

	i := 0
	for ; i < k && len(leftHeap) > 0 && len(rightHeap) > 0; i++ {
		left := float64(leftHeap[0]) + median
		right := float64(rightHeap[0]) - median

		if left <= right {
			closest[i] = -heap.Pop(&leftHeap).(int)
		} else {
			closest[i] = heap.Pop(&rightHeap).(int)
		}
	}

	for i < k && len(leftHeap) > 0 {
		closest[i] = -heap.Pop(&leftHeap).(int)
	}

	for i < k && len(rightHeap) > 0 {
		closest[i] = heap.Pop(&rightHeap).(int)
	}

	return closest
}
