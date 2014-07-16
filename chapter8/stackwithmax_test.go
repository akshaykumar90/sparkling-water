package chapter8

import "testing"

func TestStackWithMax(t *testing.T) {
	var x, m int
	st := StackWithMax{}

	st.Push(1)
	st.Push(2)
	st.Push(1)
	st.Push(5)
	st.Push(3)

	x, _ = st.Pop()
	if x != 3 {
		t.Fatalf("x = %d, want 3", x)
	}

	m, _ = st.Max()
	if m != 5 {
		t.Fatalf("m = %d, want 5", m)
	}

	st.Pop()

	m, _ = st.Max()
	if m != 2 {
		t.Fatalf("m = %d, want 2", m)
	}

	st.Push(6)
	m, _ = st.Max()
	if m != 6 {
		t.Fatalf("m = %d, want 6", m)
	}

	st.Pop()
	st.Pop()
	st.Pop()
	st.Pop()

	if _, ok := st.Max(); ok == nil {
		t.Fatalf("ok = nil, want empty stack error")
	}

}
