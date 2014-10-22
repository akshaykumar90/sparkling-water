package chapter16

import "testing"

var D = map[string]bool{
	"bat":  true,
	"cat":  true,
	"mat":  true,
	"hat":  true,
	"cot":  true,
	"door": true,
	"go":   true,
	"cut":  true,
	"put":  true,
	"pun":  true,
}

var transformStringToOtherTests = []struct {
	s        string
	t        string
	expected int
}{
	{"bat", "hat", 1},
	{"bat", "cot", 2},
	{"bat", "door", -1},
	{"bat", "go", -1},
	{"bat", "bat", 0},
	{"door", "go", -1},
	{"mat", "pun", 4},
}

func TestTransformStringToOther(t *testing.T) {
	for _, tt := range transformStringToOtherTests {
		actual := TransformStringToOther(D, tt.s, tt.t)
		if actual != tt.expected {
			t.Errorf("TransformStringToOther(%q, %q): expected %d, actual %d",
				tt.s, tt.t, tt.expected, actual)
		}
	}
}
