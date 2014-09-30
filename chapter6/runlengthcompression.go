// Problem 6.18

package chapter6

import (
	"fmt"
	"strings"
	"unicode"
)

func Encoding(s string) string {
	runes := []rune(s)

	var out string

	cnt := 1

	for i := 1; i < len(runes); i++ {
		if runes[i] != runes[i-1] {
			out += fmt.Sprintf("%d%c", cnt, runes[i-1])
			cnt = 1
		} else {
			cnt++
		}
	}

	out += fmt.Sprintf("%d%c", cnt, runes[len(runes)-1])

	return out
}

func Decoding(s string) string {
	var out string

	cnt := 0

	for _, r := range s {
		if unicode.IsDigit(r) {
			cnt = cnt*10 + int(r-'0')
		} else {
			out += strings.Repeat(string(r), cnt)
			cnt = 0
		}
	}

	return out
}
