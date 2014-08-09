// problem 5.7

package chapter5

import "unicode"

func ConvertBase(repr string, b1, b2 int) string {
	num := 0

	for _, ch := range repr {
		num = num * b1
		if unicode.IsDigit(ch) {
			num += int(ch - '0')
		} else {
			num += int(ch - 'A' + 10)
		}
	}

	slice := make([]rune, 0)

	for num > 0 {
		rem := rune(num % b2)
		if rem < 10 {
			slice = append(slice, '0'+rem)
		} else {
			rem -= 10
			slice = append(slice, 'A'+rem)
		}

		num /= b2
	}

	sz := len(slice)
	for i := 0; i < sz/2; i++ {
		slice[i], slice[sz-i-1] = slice[sz-i-1], slice[i]
	}

	if sz == 0 {
		return "0"
	} else {
		return string(slice)
	}
}
