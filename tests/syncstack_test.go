package utils_test

import (
	"sync"
	"testing"

	"github.com/k0pernicus/gostack/consts"

	"github.com/k0pernicus/gostack"
)

func TestSyncLength(t *testing.T) {
	s := gostack.NewSyncStack(consts.DefaultMaximumStackSize)
	l := s.Length()
	if l != 0 {
		t.Errorf("The stack return length is incorrect, got: %d, want %d.", l, 0)
	}
	var i uint
	var wg sync.WaitGroup
	for i = 0; i < s.Cap(); i++ {
		wg.Add(1)
		go func(i uint) {
			defer wg.Done()
			s.Push(i)
		}(i)
	}
	wg.Wait()
	l = s.Length()
	if l != consts.DefaultMaximumStackSize {
		t.Errorf("The stack return length is incorrect, got: %d, want %d.", l, consts.DefaultMaximumStackSize)
	}
}

func TestSyncPop(t *testing.T) {
	s := gostack.NewSyncStack(consts.DefaultMaximumStackSize)
	var wg sync.WaitGroup
	var i uint
	for i = 0; i < s.Cap(); i++ {
		wg.Add(1)
		go func(i uint) {
			defer wg.Done()
			s.Push(i)
		}(i)
	}
	wg.Wait()
	for i = 0; i < s.Cap(); i++ {
		wg.Add(1)
		go func(i uint) {
			defer wg.Done()
			s.Pop()
		}(i)
	}
	wg.Wait()
	if s.Length() != 0 {
		t.Errorf("The stack return value is incorrect, got: %d, want %d.", s.Length(), 0)
	}
}

func Benchmark1024SyncPop(t *testing.B) {
	for n := 0; n < t.N; n++ {
		s := gostack.NewSyncStack(consts.DefaultMaximumStackSize)
		var wg sync.WaitGroup
		var i uint
		for i = 0; i < s.Cap(); i++ {
			wg.Add(1)
			go func(i uint) {
				defer wg.Done()
				s.Push(i)
			}(i)
		}
		wg.Wait()
		for i = 0; i < s.Cap(); i++ {
			wg.Add(1)
			go func(i uint) {
				defer wg.Done()
				s.Pop()
			}(i)
		}
		wg.Wait()
	}
}

// func TestMaximumStack(t *testing.T) {
// 	s := gostack.NewSyncStack(10)
// 	for i := 0; i < utils.StackMaximumStackSize; i++ {
// 		s.Push(i)
// 	}
// 	c := s.Cap()
// 	if c != utils.StackMaximumStackSize || s.Length() != utils.StackMaximumStackSize {
// 		t.Errorf("The stack capacity is incorrect, got: %d, want %d.", c, utils.StackMaximumStackSize)
// 	}
// 	// The stack canno't expend after StackMaximumStackSize
// 	for i := 0; i < utils.StackMaximumStackSize; i++ {
// 		s.Push(i)
// 	}
// 	c = s.Cap()
// 	if c != utils.StackMaximumStackSize || s.Length() != utils.StackMaximumStackSize {
// 		t.Errorf("The stack capacity should not have widened, got: %d, want %d.", c, utils.StackMaximumStackSize)
// 	}
// }
