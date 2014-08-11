// problem 6.1

package chapter6

func DutchNationalFlag(arr []int, p int) {
	pivot := arr[p]
	smaller, equal, larger := 0, 0, len(arr)-1

	for equal <= larger {
		switch {
		case arr[equal] < pivot:
			arr[smaller], arr[equal] = arr[equal], arr[smaller]
			smaller++
			equal++
		case arr[equal] == pivot:
			equal++
		case arr[equal] > pivot:
			arr[larger], arr[equal] = arr[equal], arr[larger]
			larger--
		}
	}
}
