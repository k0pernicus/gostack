package gostack

import (
	"github.com/k0pernicus/gostack/consts"
	"github.com/k0pernicus/gostack/errors"
)

type Stack struct {
	capacity    uint
	cLength     uint
	content     []interface{}
	maxCapacity uint
}

func NewStack(capacity uint) Stack {
	maxCapacity := consts.DefaultMaximumStackSize
	if capacity > maxCapacity {
		maxCapacity = capacity
	}
	// Avoid 'append' call for each Push, in declaring size == capacity
	return Stack{
		capacity:    capacity,
		cLength:     0,
		content:     make([]interface{}, capacity, capacity),
		maxCapacity: maxCapacity}
}

func (s *Stack) Push(v interface{}) error {
	if s.cLength >= s.capacity {
		if s.capacity >= s.maxCapacity {
			return errors.MaximumStackCapacityExceededError
		}
		s.capacity = s.capacity * 2
		if s.capacity > s.maxCapacity {
			s.capacity = s.maxCapacity
		}
		newContent := make([]interface{}, s.capacity, s.capacity)
		copy(newContent, s.content)
		s.content = newContent
	}
	s.content[s.cLength] = v
	s.cLength++
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.cLength == 0 {
		return nil, errors.NoMoreContentError
	}
	s.cLength--
	value := s.content[s.cLength]
	s.content = s.content[:s.cLength]
	return value, nil
}

func (s *Stack) MaxCap(maxCap uint) error {
	if maxCap == s.capacity {
		return nil
	}
	if maxCap > s.capacity && maxCap > s.cLength {
		return errors.MinCapacityError
	}
	s.maxCapacity = maxCap
	return nil
}

func (s *Stack) Length() uint {
	return s.cLength
}

func (s *Stack) Cap() uint {
	return s.capacity
}

func (s *Stack) Empty() bool {
	return s.cLength == 0
}
