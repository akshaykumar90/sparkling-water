// Problem 10.10

package chapter10

func CompareKthLargestInHeap(heap []int, k, x int) int {
	larger, equal := 0, 0

	var helper func(i int)
	helper = func(i int) {
		if larger >= k || i >= len(heap) || heap[i] < x {
			return
		}

		if heap[i] == x {
			equal++
			if equal >= k {
				return
			}
		} else {
			larger++
		}

		helper(2*i + 1)
		helper(2*i + 2)
	}

	helper(0)

	if larger >= k {
		return 1
	} else if larger+equal >= k {
		return 0
	} else {
		return -1
	}
}
