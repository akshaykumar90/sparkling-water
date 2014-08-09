// problem 5.6

package chapter5

import "errors"

func StringToInt(s string) (int, error) {
	if len(s) == 0 {
		return 0, errors.New("StringToInt: invalid string")
	}

	mult, num := 1, 0

	for i, ch := range s {
		switch {
		case ch == '-':
			if i == 0 {
				mult = -1
			} else {
				return 0, errors.New("StringToInt: invalid string")
			}
		case ch > '9' || ch < '0':
			return 0, errors.New("StringToInt: invalid string")
		default:
			num = num*10 + int(ch-'0')
		}
	}

	if mult == -1 && len(s) == 1 {
		return 0, errors.New("StringToInt: invalid string")
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
