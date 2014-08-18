// problem 8.6

package chapter8

func ViewSunset(hs []int) IntStack {
	view := make(IntStack, 0)

	for _, h := range hs {
		for len(view) > 0 && view[len(view)-1] <= h {
			view.Pop()
		}
		view.Push(h)
	}

	return view
}
