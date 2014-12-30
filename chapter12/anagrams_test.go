package chapter12

import (
	"sort"
	"testing"
)

func TestAnagrams(t *testing.T) {
	words := []string{"debit card", "bad credit", "the morse code", "here come dots", "the eyes", "they see", "THL"}

	expected := map[string][]string{
		"debit card":     {"bad credit", "debit card"},
		"bad credit":     {"bad credit", "debit card"},
		"the morse code": {"here come dots", "the morse code"},
		"here come dots": {"here come dots", "the morse code"},
		"the eyes":       {"the eyes", "they see"},
		"they see":       {"the eyes", "they see"},
	}

	actual := Anagrams(words)

	if len(actual) != 3 {
		t.Fatalf("Anagrams: expected length %d, actual length %d", 3, len(actual))
	}

	for _, subset := range actual {
		fst := subset[0]
		if _, ok := expected[fst]; !ok {
			t.Fatalf("Anagrams: Unexpected %q, not present in any anagrams subset.", fst)
		}

		ess := expected[fst]
		if len(subset) != len(ess) {
			t.Fatalf("Anagrams: subset %q: expected length %d, actual length %d", fst, len(ess), len(subset))
		}

		sort.Strings(subset)
		for i, v := range ess {
			if v != subset[i] {
				t.Fatalf("Anagrams: subset %q: at %d: expected %q, actual %q", fst, i, v, subset[i])
			}
		}
	}
}
