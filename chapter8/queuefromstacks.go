// Problem 8.12

package chapter8

import "errors"

type IntStack []int

func (st *IntStack) Push(n int) {
	*st = append(*st, n)
}

func (st *IntStack) Pop() (int, error) {
	if len(*st) == 0 {
		return 0, errors.New("Pop: empty stack")
	} else {
		var x int
		x, *st = (*st)[len(*st)-1], (*st)[:len(*st)-1]
		return x, nil
	}
}

type StackedQueue struct {
	head, tail IntStack
}

func (sq *StackedQueue) Enqueue(n int) {
	sq.head.Push(n)
}

func (sq *StackedQueue) Dequeue() (int, error) {
	if len((*sq).tail) > 0 {
		return sq.tail.Pop()
	} else if len((*sq).head) > 0 {
		for len((*sq).head) > 0 {
			n, _ := sq.head.Pop()
			sq.tail.Push(n)
		}
		return sq.tail.Pop()
	} else {
		return 0, errors.New("Dequeue: empty queue")
	}
}

func NewStackedQueue() *StackedQueue {
	head := make(IntStack, 0)
	tail := make(IntStack, 0)
	return &StackedQueue{head, tail}
}
