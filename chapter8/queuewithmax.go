// problem 8.13

package chapter8

import (
	"container/list"
	"errors"
)

type QueueWithMax struct {
	q, d *list.List
}

func NewQueueWithMax() *QueueWithMax {
	return &QueueWithMax{list.New(), list.New()}
}

func (qwm *QueueWithMax) Enqueue(n int) {
	qwm.q.PushBack(n)
	for e := qwm.d.Back(); e != nil; {
		if e.Value.(int) < n {
			tmp := e
			e = e.Prev()
			qwm.d.Remove(tmp)
		} else {
			e = e.Prev()
		}
	}
	qwm.d.PushBack(n)
}

func (qwm *QueueWithMax) Dequeue() (int, error) {
	if elem := qwm.q.Front(); elem != nil {
		n := qwm.q.Remove(elem).(int)
		if maxElem := qwm.d.Front(); n == maxElem.Value.(int) {
			qwm.d.Remove(maxElem)
		}
		return n, nil
	} else {
		return 0, errors.New("Dequeue: empty queue")
	}
}

func (qwm *QueueWithMax) Max() (int, error) {
	if maxElem := qwm.d.Front(); maxElem != nil {
		return maxElem.Value.(int), nil
	} else {
		return 0, errors.New("Max: empty queue")
	}
}
