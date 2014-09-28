package chapter7

import "testing"

var palindromeLinkedListTests = []struct {
	head     *Element
	expected bool
}{
	{
		&Element{1, &Element{2, &Element{3, &Element{4, &Element{5, nil}}}}},
		false,
	},
	{
		&Element{1, &Element{2, &Element{3, &Element{2, &Element{1, nil}}}}},
		true,
	},
	{
		&Element{1, &Element{2, &Element{3, &Element{2, &Element{1, &Element{1, nil}}}}}},
		false,
	},
	{
		&Element{1, &Element{2, &Element{3, &Element{3, &Element{2, &Element{1, nil}}}}}},
		true,
	},
}

func TestPalindromeLinkedList(t *testing.T) {
	for _, tt := range palindromeLinkedListTests {
		actual := PalindromeLinkedList(tt.head)
		if actual != tt.expected {
			t.Errorf("PalindromeLinkedList(%v): expected %t, actual %t",
				tt.head, tt.expected, actual)
		}
	}
}
