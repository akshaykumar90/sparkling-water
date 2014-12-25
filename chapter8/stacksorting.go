// Problem 8.7

package chapter8

func insert(st *IntStack, x int) {
	if len(*st) == 0 || x >= (*st)[len(*st)-1] {
		st.Push(x)
	} else {
		top, _ := st.Pop()
		insert(st, x)
		st.Push(top)
	}
}

func StackSorting(st *IntStack) {
	if len(*st) > 0 {
		top, _ := st.Pop()
		StackSorting(st)
		insert(st, top)
	}
}
