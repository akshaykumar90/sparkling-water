package chapter13

import (
	"fmt"
	"testing"
)

func TestCountOccurrencesInSentence(t *testing.T) {
	s := "bcdacebe"
	fmt.Printf("CountOccurrencesInSentence(%q)\n", s)
	CountOccurrencesInSentence(s)
	fmt.Println()
}
