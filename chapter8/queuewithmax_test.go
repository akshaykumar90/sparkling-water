package chapter8

import "testing"

func TestQueueWithMax(t *testing.T) {
	qwm := NewQueueWithMax()

	var n, m int

	qwm.Enqueue(3)
	qwm.Enqueue(1)
	qwm.Enqueue(3)
	qwm.Enqueue(2)
	qwm.Enqueue(0)
	qwm.Enqueue(1)

	m, _ = qwm.Max()
	if m != 3 {
		t.Fatalf("m = %d, want 3", m)
	}

	n, _ = qwm.Dequeue()
	if n != 3 {
		t.Fatalf("n = %d, want 3", n)
	}

	n, _ = qwm.Dequeue()
	if n != 1 {
		t.Fatalf("n = %d, want 1", n)
	}

	m, _ = qwm.Max()
	if m != 3 {
		t.Fatalf("m = %d, want 3", m)
	}

	qwm.Enqueue(2)
	qwm.Enqueue(4)
	m, _ = qwm.Max()
	if m != 4 {
		t.Fatalf("m = %d, want 4", m)
	}

	n, _ = qwm.Dequeue()
	if n != 3 {
		t.Fatalf("n = %d, want 3", n)
	}

	qwm.Enqueue(4)
	m, _ = qwm.Max()
	if m != 4 {
		t.Fatalf("m = %d, want 4", m)
	}
}
