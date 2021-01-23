package stack

import "errors"

type uint64Stack struct {
	datas []uint64
}

func (s *uint64Stack) Push(n uint64) {
	s.datas = append(s.datas, n)
}

func (s *uint64Stack) Pop() (uint64, error) {
	if len(s.datas) == 0 {
		return 0, errors.New("try pop from empty stack")
	}

	top := s.datas[len(s.datas)-1]
	s.datas = s.datas[:len(s.datas)-1]
	return top, nil
}
