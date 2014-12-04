package chapter15

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

var text = `I have inserted a large number of new examples from the papers for the 
Mathematical Tripos during the last twenty years, which should be useful 
to Cambridge students.`

var width = 36

func TestPrettyPrintingCross(t *testing.T) {
	for w := 18; w < 64; w++ {
		_, m1 := PrettyPrinting(text, w)
		_, m2 := PrettyPrinting2(text, w)

		if m1 != m2 {
			t.Errorf("PrettyPrinting(%d) != PrettyPrinting2(%d): %d, %d", w, w, m1, m2)
		}
	}
}

func TestPrettyPrinting(t *testing.T) {
	b, min := PrettyPrinting(text, width)

	words := strings.Fields(text)

	fmt.Println("TestPrettyPrinting:")

	for k := 0; k < len(words); k = b[k] {
		line := strings.Join(words[k:b[k]], " ")
		fmt.Printf("%-*s%d\n", width, line, width-utf8.RuneCountInString(line))
	}

	fmt.Println(min)

	fmt.Println()
}

func TestPrettyPrinting2(t *testing.T) {
	b, min := PrettyPrinting2(text, width)

	words := strings.Fields(text)

	fmt.Println("TestPrettyPrinting2:")

	lines := make([]string, 0)

	for k := len(words); k > 0; k = b[k] {
		lines = append(lines, strings.Join(words[b[k]:k], " "))
	}

	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		fmt.Printf("%-*s%d\n", width, line, width-utf8.RuneCountInString(line))
	}

	fmt.Println(min)

	fmt.Println()
}
