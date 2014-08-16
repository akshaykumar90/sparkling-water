package chapter7

import "testing"

var removeKthLastListTests = []struct {
	list     *Element
	k        int
	expected *Element
}{
	{
		&Element{1, &Element{2, &Element{3, &Element{4, &Element{5, nil}}}}},
		3,
		&Element{1, &Element{2, &Element{4, &Element{5, nil}}}},
	},
	{
		&Element{1, &Element{2, &Element{3, &Element{4, &Element{5, nil}}}}},
		5,
		&Element{2, &Element{3, &Element{4, &Element{5, nil}}}},
	},
	{
		&Element{1, &Element{2, &Element{3, &Element{4, &Element{5, nil}}}}},
		1,
		&Element{1, &Element{2, &Element{3, &Element{4, nil}}}},
	},
	{
		&Element{1, nil},
		1,
		nil,
	},
}

func TestRemoveKthLastList(t *testing.T) {
	for _, tt := range removeKthLastListTests {
		if actual, err := RemoveKthLastList(tt.list, tt.k); err == nil {
			assertEqualLists(t, "RemoveKthLastList", tt.expected, actual)
		} else {
			t.Errorf("RemoveKthLastList(%v, %d): unexpected error %s",
				tt.list, tt.k, err)
		}
	}
}

var invalidRemoveKthLastListTests = []struct {
	list *Element
	k    int
}{
	{&Element{1, &Element{2, &Element{3, &Element{4, &Element{5, nil}}}}}, 6},
	{&Element{1, nil}, 2},
}

func TestRemoveKthLastListInvalid(t *testing.T) {
	for _, tt := range invalidRemoveKthLastListTests {
		if actual, err := RemoveKthLastList(tt.list, tt.k); err == nil {
			t.Errorf("RemoveKthLastList(%v, %d): expected error, actual %v",
				tt.list, tt.k, actual)
		}
	}
}
