// problem 8.1

package chapter8

import "errors"

type StackWithMax struct {
	elements []int
	max      []int
}

func (s *StackWithMax) Push(x int) {
	if len(s.elements) == 0 {
		s.max = append(s.max, x)
		s.elements = append(s.elements, x)
	} else {
		if x < s.elements[len(s.elements)-1] {
			s.max = append(s.max, s.elements[len(s.elements)-1])
		} else {
			s.max = append(s.max, x)
		}
		s.elements = append(s.elements, x)
	}
}

func (s *StackWithMax) Pop() (int, error) {
	es := s.elements

	if len(es) == 0 {
		return 0, errors.New("Pop: empty stack")
	}

	x := es[len(es)-1]
	s.elements = es[:len(es)-1]
	s.max = s.max[:len(s.max)-1]

	return x, nil
}

func (s *StackWithMax) Max() (int, error) {
	if len(s.max) == 0 {
		return 0, errors.New("Max: empty stack")
	}

	return s.max[len(s.max)-1], nil
}
