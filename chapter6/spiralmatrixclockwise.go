// problem 6.15

package chapter6

func spiralMatrixPerimeter(arr [][]int, st int) []int {
	side := len(arr) - 2*st

	if side == 1 {
		return []int{arr[st][st]}
	} else {
		seq := make([]int, 4*(side-1))
		for i := 0; i < side-1; i++ {
			seq[i] = arr[st][st+i]
			seq[(side-1)+i] = arr[st+i][st+side-1]
			seq[2*(side-1)+i] = arr[st+side-1][st+side-1-i]
			seq[3*(side-1)+i] = arr[st+side-1-i][st]
		}
		return seq
	}
}

func SpiralMatrixClockwise(arr [][]int) []int {
	res := make([]int, 0)
	for i := 0; i < (len(arr)+1)/2; i++ {
		res = append(res, spiralMatrixPerimeter(arr, i)...)
	}
	return res
}
