// problem 6.19

package chapter6

func reverse(str []rune) {
	sz := len(str)
	for i := 0; i < sz/2; i++ {
		str[i], str[sz-i-1] = str[sz-i-1], str[i]
	}
}

func ReverseWords(str string) string {
	strSlice := []rune(str)

	sz := len(strSlice)

	reverse(strSlice)

	prev := 0
	for i := 0; i < sz; i++ {
		switch {
		case strSlice[i] == ' ':
			reverse(strSlice[prev:i])
			prev = i
		case strSlice[prev] == ' ':
			prev = i
		}
	}

	reverse(strSlice[prev:])

	return string(strSlice)
}
