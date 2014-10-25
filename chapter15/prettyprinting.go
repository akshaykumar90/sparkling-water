// Problem 15.13

package chapter15

import (
	"math"
	"strings"
	"unicode/utf8"
)

func PrettyPrinting(text string, width int) ([]int, int) {
	words := strings.Fields(text)

	wordLengths := make([]int, len(words)+1)

	for i, w := range words {
		wordLengths[i] = utf8.RuneCountInString(w)
	}

	m, b := make([]int, len(words)+1), make([]int, len(words)+1)

	for i := len(words) - 1; i >= 0; i-- {
		minMess := math.MaxInt64

		j, ws := i+1, width-wordLengths[i]
		for ; j <= len(words) && ws > 0; j++ {
			mess := 1<<uint(ws) + m[j]
			if mess < minMess {
				minMess = mess
				b[i] = j
			}
			ws -= wordLengths[j] + 1
		}

		m[i] = minMess
	}

	return b, m[0]
}

func PrettyPrinting2(text string, width int) ([]int, int) {
	words := strings.Fields(text)

	wordLengths := make([]int, len(words)+1)

	for i, w := range words {
		wordLengths[i+1] = utf8.RuneCountInString(w)
	}

	m, b := make([]int, len(words)+1), make([]int, len(words)+1)

	for i := 1; i <= len(words); i++ {
		minMess := math.MaxInt64

		j, ws := i-1, width-wordLengths[i]
		for ; j >= 0 && ws > 0; j-- {
			mess := 1<<uint(ws) + m[j]
			if mess < minMess {
				minMess = mess
				b[i] = j
			}
			ws -= wordLengths[j] + 1
		}

		m[i] = minMess
	}

	return b, m[len(words)]
}
