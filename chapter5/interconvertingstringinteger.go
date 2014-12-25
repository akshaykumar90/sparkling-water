// Problem 5.6

package chapter5

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("StringToInt: invalid string")

func StringToInt(s string) (int, error) {
	if len(s) == 0 || s == "-" {
		return 0, ErrInvalidString
	}

	mult, num := 1, 0

	fst, size := utf8.DecodeRuneInString(s)
	if fst == '-' {
		mult = -1
		s = s[size:]
	}

	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			return 0, ErrInvalidString
		}
		num = num*10 + int(ch-'0')
	}

	return mult * num, nil
}

func IntToString(x int) string {
	if x == 0 {
		return "0"
	}

	slice := make([]rune, 0)

	isNegative := false

	if x < 0 {
		x = -x
		isNegative = true
	}

	for x > 0 {
		slice = append(slice, '0'+rune(x%10))
		x = x / 10
	}

	if isNegative {
		slice = append(slice, '-')
	}

	n := len(slice)
	for i := 0; i < n/2; i++ {
		slice[i], slice[n-i-1] = slice[n-i-1], slice[i]
	}

	return string(slice)
}
