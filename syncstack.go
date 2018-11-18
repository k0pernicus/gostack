package gostack

import (
	"sync"
)

type SyncStack struct {
	mux   sync.Mutex
	stack Stack
}

func NewSyncStack(capacity uint) SyncStack {
	return SyncStack{
		stack: NewStack(capacity)}
}

func (s *SyncStack) Push(v interface{}) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.stack.Push(v)
}

func (s *SyncStack) Pop() (interface{}, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.stack.Pop()
}

func (s *SyncStack) MaxCap(maxCap uint) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.stack.MaxCap(maxCap)
}

func (s *SyncStack) Length() uint {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.stack.Length()
}

func (s *SyncStack) Cap() uint {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.stack.Cap()
}

func (s *SyncStack) Empty() bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.stack.Empty()
}
