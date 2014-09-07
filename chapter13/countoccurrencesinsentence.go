// problem 13.7

package chapter13

import (
	"fmt"
	"sort"
)

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func CountOccurrencesInSentence(s string) {
	runes := []rune(s)

	sort.Sort(RuneSlice(runes))

	cnt := 1
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			cnt++
		} else {
			fmt.Printf("(%c,%d),", runes[i-1], cnt)
			cnt = 1
		}
	}
	fmt.Printf("(%c,%d)\n", runes[len(runes)-1], cnt)
}
