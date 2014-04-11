package chapter3

import (
	"errors"
)

type Stack struct {
	m     []interface{}
	index int
}

var EmptyStackError error = errors.New("Empty Stack")

func NewStack() *Stack {
	return &Stack{
		m:     make([]interface{}, 1000),
		index: 0,
	}
}

func (s *Stack) expand() {
	new_memory := make([]interface{}, len(s.m)*2)
	copy(new_memory, s.m)
	s.m = new_memory
}

func (s *Stack) Push(item interface{}) {
	if (s.index + 1) == len(s.m) {
		s.expand()
	}
	s.index += 1
	s.m[s.index] = item
}

func (s *Stack) Pop() (interface{}, error) {
	if s.index == 0 {
		return nil, EmptyStackError
	}
	item := s.m[s.index]
	s.index -= 1
	return item, nil
}
