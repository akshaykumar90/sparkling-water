// Problem 6.22

package chapter6

var M = []string{
	"0",
	"1",
	"ABC",
	"DEF",
	"GHI",
	"JKL",
	"MNO",
	"PQRS",
	"TUV",
	"WXYZ",
}

func PhoneMnemonic(phone []int) []string {
	result := make([]string, 0)

	partial := make([]rune, len(phone))

	var helper func(int)
	helper = func(idx int) {
		if idx == len(phone) {
			result = append(result, string(partial))
		} else {
			for _, r := range M[phone[idx]] {
				partial[idx] = r
				helper(idx + 1)
			}
		}
	}

	helper(0)

	return result
}
