package chapter10

import (
	"github.com/akshaykumar90/sparkling-water/common"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestKThLargestElement(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	stream := rand.Perm(100)
	k := 15

	expected := make([]int, 0, 100)

	temp := make([]int, 100)
	copy(temp, stream)
	for i := range temp {
		arr := temp[:i+1]
		sort.Ints(arr)

		if i+1 <= k {
			expected = append(expected, arr[0])
		} else {
			expected = append(expected, arr[len(arr)-k])
		}
	}

	actual := KThLargestElement(stream, k)

	common.AssertIntsAreEqual(t, "KThLargestElement", expected, actual)
}
