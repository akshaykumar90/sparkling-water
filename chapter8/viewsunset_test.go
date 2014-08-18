package chapter8

import (
	"math/rand"
	"testing"
	"time"
)

func TestViewSunset(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	hs := make([]int, 100)

	for i := range hs {
		hs[i] = rand.Intn(200) + 1
	}

	res := ViewSunset(hs)

	for i := 1; i < len(res); i++ {
		if res[i-1] <= res[i] {
			t.Fatalf("res[%d]:%d <= res[%d]:%d", i-1, res[i-1], i, res[i])
		}
	}
}
