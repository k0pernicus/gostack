package utils_test

import (
	"testing"

	"github.com/k0pernicus/gostack/consts"

	"github.com/k0pernicus/gostack"
	"github.com/k0pernicus/mercure/utils"
)

func TestCapacity(t *testing.T) {
	s := gostack.NewStack(1)
	c := s.Cap()
	if c != 1 {
		t.Errorf("The stack return capacity is incorrect, got: %d, want %d.", c, 1)
	}
}

func TestLength(t *testing.T) {
	s := gostack.NewStack(1)
	l := s.Length()
	if l != 0 {
		t.Errorf("The stack return length is incorrect, got: %d, want %d.", l, 0)
	}
	s.Push(0)
	l = s.Length()
	if l != 1 {
		t.Errorf("The stack return length is incorrect, got: %d, want %d.", l, 1)
	}
	s.Pop()
	l = s.Length()
	if l != 0 {
		t.Errorf("The final stack return length is incorrect, got: %d, want %d.", l, 0)
	}
}

func TestPop(t *testing.T) {
	s := gostack.NewStack(1)
	s.Push(0)
	v, _ := s.Pop()
	if v != 0 {
		t.Errorf("The stack return value is incorrect, got: %d, want %d.", v, 0)
	}
}

func Benchmark1024Pop(t *testing.B) {
	for n := 0; n < t.N; n++ {
		s := gostack.NewStack(consts.DefaultMaximumStackSize)
		var i uint
		for i = 0; i < s.Cap(); i++ {
			s.Push(i)
		}
		for i = 0; i < s.Cap(); i++ {
			s.Pop()
		}
	}
}

func TestMaximumStack(t *testing.T) {
	s := gostack.NewStack(10)
	for i := 0; i < utils.StackMaximumStackSize; i++ {
		s.Push(i)
	}
	c := s.Cap()
	if c != utils.StackMaximumStackSize || s.Length() != utils.StackMaximumStackSize {
		t.Errorf("The stack capacity is incorrect, got: %d, want %d.", c, utils.StackMaximumStackSize)
	}
	// The stack canno't expend after StackMaximumStackSize
	for i := 0; i < utils.StackMaximumStackSize; i++ {
		s.Push(i)
	}
	c = s.Cap()
	if c != utils.StackMaximumStackSize || s.Length() != utils.StackMaximumStackSize {
		t.Errorf("The stack capacity should not have widened, got: %d, want %d.", c, utils.StackMaximumStackSize)
	}
}
