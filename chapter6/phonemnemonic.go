// problem 6.22

package chapter6

import "fmt"

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

func phoneMnemonicHelper(phone []int, M []string, partial []rune, idx int) {
	if idx == len(phone) {
		fmt.Println(string(partial))
	} else {
		for _, r := range M[phone[idx]] {
			partial[idx] = r
			phoneMnemonicHelper(phone, M, partial, idx+1)
		}
	}
}

func PhoneMnemonic(phone []int) {
	partial := make([]rune, len(phone))
	phoneMnemonicHelper(phone, M, partial, 0)
}
