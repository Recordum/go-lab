package stack

import (
	. "go-lab/data-structure/common"
)

type Stack struct {
	top *Elem
	len int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Len() int {
	return s.len
}

func (s* Stack) Push(value interface{}) {
	elem := &Elem{Value: value}
	if s.top == nil {
		s.top = elem
	} else {
		elem.Next = s.top
		s.top.Prev = elem
		s.top = elem
	}
	s.len++
}

