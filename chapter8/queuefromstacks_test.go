package chapter8

import "testing"

func TestStackedQueue(t *testing.T) {
	sq := NewStackedQueue()

	var n int

	sq.Enqueue(1)
	sq.Enqueue(2)
	sq.Enqueue(3)

	n, _ = sq.Dequeue()
	if n != 1 {
		t.Fatalf("n = %d, want 1", n)
	}

	sq.Enqueue(4)
	sq.Enqueue(5)

	n, _ = sq.Dequeue()
	if n != 2 {
		t.Fatalf("n = %d, want 2", n)
	}

	sq.Dequeue()

	n, _ = sq.Dequeue()
	if n != 4 {
		t.Fatalf("n = %d, want 4", n)
	}

	sq.Dequeue()

	if _, ok := sq.Dequeue(); ok == nil {
		t.Fatalf("ok = nil, want empty queue error")
	}
}
