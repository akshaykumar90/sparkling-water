package chapter9

import "testing"

func TestBinaryTreeLock(t *testing.T) {
	elems := make([]BinaryTreeLock, 5)

	elems[0].Left = &elems[1]
	elems[1].Parent = &elems[0]

	elems[0].Right = &elems[2]
	elems[2].Parent = &elems[0]

	elems[1].Left = &elems[3]
	elems[3].Parent = &elems[1]

	elems[1].Right = &elems[4]
	elems[4].Parent = &elems[1]

	// ---

	if elems[0].isLock() {
		t.Fatalf("isLock: %t, want %t", true, false)
	}

	// ---

	elems[0].lock()

	if !elems[0].isLock() {
		t.Fatalf("isLock: %t, want %t", false, true)
	}

	elems[0].unlock()

	// ---

	elems[1].lock()
	elems[0].lock()

	if elems[0].isLock() {
		t.Fatalf("isLock: %t, want %t", true, false)
	}

	// ---

	elems[2].lock()

	if !elems[2].isLock() {
		t.Fatalf("isLock: %t, want %t", false, true)
	}
}
