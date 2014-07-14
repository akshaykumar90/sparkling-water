package chapter6

import (
	"fmt"
	"testing"
)

func TestPhoneMnemonic(t *testing.T) {
	var phoneMnemonicTests = [][]int{
		{2},
		{2,3},
		// {2,2,7,6,6,9,6},
	}

	for _, tt := range phoneMnemonicTests {
		fmt.Println(tt)
		PhoneMnemonic(tt)
		fmt.Println()
	}
}

