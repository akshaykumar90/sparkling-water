package chapter15

import "testing"

func TestWordBreaking(t *testing.T) {
	dict := map[string]bool{
		"a":      true,
		"an":     true,
		"and":    true,
		"at":     true,
		"bat":    true,
		"bath":   true,
		"be":     true,
		"bed":    true,
		"beyond": true,
		"hand":   true,
		"on":     true,
		"than":   true,
	}

	s := "bedbathandbeyond"

	actual := WordBreaking(dict, s)

	if true != actual {
		t.Errorf("WordBreaking %q: expected %d, actual %d", s, true, actual)
	}
}
