package chapter16

import "testing"

func TestMinimumDelaySchedule1(t *testing.T) {
	a := Task{1, 1, []Task{}}
	b := Task{2, 1, []Task{a}}
	c := Task{3, 2, []Task{a}}
	d := Task{4, 3, []Task{b, c}}

	tasks := []Task{a, b, c, d}

	if actual, expected := MinimumDelaySchedule(tasks), 6; actual != expected {
		t.Errorf("MinimumDelaySchedule(%v): expected %d, actual %d", tasks, expected, actual)
	}
}

func TestMinimumDelaySchedule2(t *testing.T) {
	a := Task{1, 1, []Task{}}
	b := Task{2, 1, []Task{a}}
	c := Task{3, 2, []Task{a, b}}
	d := Task{4, 3, []Task{b, c}}

	tasks := []Task{a, b, c, d}

	if actual, expected := MinimumDelaySchedule(tasks), 7; actual != expected {
		t.Errorf("MinimumDelaySchedule(%v): expected %d, actual %d", tasks, expected, actual)
	}
}

func TestMinimumDelaySchedule3(t *testing.T) {
	a := Task{1, 1, []Task{}}
	b := Task{2, 2, []Task{a}}
	c := Task{3, 3, []Task{b}}

	a.p = append(a.p, c)

	tasks := []Task{a, b, c}

	if actual, expected := MinimumDelaySchedule(tasks), -1; actual != expected {
		t.Errorf("MinimumDelaySchedule(%v): expected %d, actual %d", tasks, expected, actual)
	}
}

func TestMinimumDelaySchedule4(t *testing.T) {
	a := Task{1, 1, []Task{}}

	tasks := []Task{a}

	if actual, expected := MinimumDelaySchedule(tasks), 1; actual != expected {
		t.Errorf("MinimumDelaySchedule(%v): expected %d, actual %d", tasks, expected, actual)
	}
}

func TestMinimumDelaySchedule5(t *testing.T) {
	a := Task{1, 1, []Task{}}
	b := Task{2, 1, []Task{a}}
	c := Task{3, 2, []Task{a}}
	d := Task{4, 3, []Task{b, c}}
	e := Task{5, 9, []Task{c}}

	tasks := []Task{a, b, c, d, e}

	if actual, expected := MinimumDelaySchedule(tasks), 12; actual != expected {
		t.Errorf("MinimumDelaySchedule(%v): expected %d, actual %d", tasks, expected, actual)
	}
}
