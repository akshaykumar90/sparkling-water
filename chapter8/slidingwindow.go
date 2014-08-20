// problem 8.14

package chapter8

import "container/list"

type TrafficElement struct {
	ts, vol int
}

func SlidingWindow(t []TrafficElement, w int) []int {
	qTs := list.New()
	qwm := NewQueueWithMax()

	max := make([]int, len(t))

	for i, te := range t {
		qwm.Enqueue(te.vol)
		qTs.PushBack(te.ts)
		for te.ts-qTs.Front().Value.(int) > w {
			qwm.Dequeue()
			qTs.Remove(qTs.Front())
		}
		m, _ := qwm.Max()
		max[i] = m
	}

	return max
}
