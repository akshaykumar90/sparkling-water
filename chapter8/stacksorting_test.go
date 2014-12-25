package chapter8

import (
	"math/rand"
	"testing"
	"time"
)

func TestStackSorting(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	st := make(IntStack, 100)

	for i := range st {
		st[i] = rand.Intn(1000) + 1
	}

	StackSorting(&st)

	for i := 1; i < len(st); i++ {
		if st[i-1] > st[i] {
			t.Fatalf("StackSorting: st[%d]:%d > st[%d]:%d", i-1, st[i-1], i, st[i])
		}
	}
}
