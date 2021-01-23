package stack

import "testing"

func TestUint64StackWorksAsStack(t *testing.T) {
	s := NewUint64Stack()

	assertSucceed := func(n uint64, err error) func(uint64, *testing.T) {
		return func(expected uint64, t *testing.T) {
			if err != nil {
				t.Error(err)
				return
			}
			if n != expected {
				t.Errorf("expected %v, but actual %v", expected, n)
			}
		}
	}

	assertFail := func(_ uint64, err error) func(*testing.T) {
		return func(t *testing.T) {
			if err == nil {
				t.Error("an error is expected, but nil")
			}
		}
	}

	assertFail(s.Pop())(t)

	s.Push(1)
	s.Push(2)
	s.Push(3)

	assertSucceed(s.Pop())(3, t)

	s.Push(4)
	assertSucceed(s.Pop())(4, t)
	assertSucceed(s.Pop())(2, t)
	assertSucceed(s.Pop())(1, t)
	assertFail(s.Pop())(t)
}
