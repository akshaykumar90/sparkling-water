package chapter12

import (
	"fmt"
	"testing"
)

func TestAnagrams(t *testing.T) {
	words := []string{"debit card", "bad credit", "the morse code", "here come dots", "the eyes", "they see", "THL"}

	fmt.Println(Anagrams(words))
	fmt.Println()

}
