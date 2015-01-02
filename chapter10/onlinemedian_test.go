package chapter10

import (
	"fmt"
	"testing"
)

func TestOnlineMedian(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(OnlineMedian(in))
	fmt.Println()
}
