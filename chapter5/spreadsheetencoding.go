// problem 5.8

package chapter5

func SpreadsheetEncoding(col string) int {
	num := 0

	for _, ch := range col {
		num = num*26 + int(ch-'A') + 1
	}

	return num
}
