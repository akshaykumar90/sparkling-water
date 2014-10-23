// Problem 16.8

package chapter16

type Task struct {
	id int
	d  int
	p  []Task
}

func MinimumDelaySchedule(tasks []Task) int {
	feasible := true

	topsort := make([]Task, 0)

	state := make(map[int]int)

	for _, t := range tasks {
		state[t.id] = UNDISCOVERED
	}

	var dfs func(Task)
	dfs = func(task Task) {
		state[task.id] = DISCOVERED
		for _, p := range task.p {
			if state[p.id] == DISCOVERED {
				feasible = false
				return
			} else if state[p.id] == UNDISCOVERED {
				dfs(p)
				if !feasible {
					return
				}
			}
		}
		state[task.id] = VISITED
		topsort = append(topsort, task)
	}

	for _, t := range tasks {
		if state[t.id] == UNDISCOVERED {
			dfs(t)
		}
		if !feasible {
			return -1
		}
	}

	timeToComplete := make(map[int]int)

	mds := 0

	for _, t := range topsort {
		currMax := 0
		for _, p := range t.p {
			if currMax < timeToComplete[p.id] {
				currMax = timeToComplete[p.id]
			}
		}

		timeToComplete[t.id] = currMax + t.d

		if mds < timeToComplete[t.id] {
			mds = timeToComplete[t.id]
		}
	}

	return mds
}
