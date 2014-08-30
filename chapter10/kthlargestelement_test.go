package chapter10

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestKThLargestElement(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	stream := rand.Perm(100)
	fmt.Println(stream)

	KThLargestElement(stream, 15)
	fmt.Println()
}
