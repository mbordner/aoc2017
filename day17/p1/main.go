package main

import "fmt"

type SpinLock struct {
	ptr int
	mem []int
	nv  int
}

func (s *SpinLock) Next(steps int) {
	s.ptr += steps
	for s.ptr >= len(s.mem) {
		s.ptr -= len(s.mem)
	}
	s.ptr++
	s.nv++
	s.mem = append(s.mem[:s.ptr], append([]int{s.nv}, s.mem[s.ptr:]...)...)
}

func (s *SpinLock) GetValue(offset int) int {
	ptr := s.ptr + offset
	for ptr < 0 {
		ptr += len(s.mem)
	}
	for ptr >= len(s.mem) {
		ptr -= len(s.mem)
	}
	return s.mem[ptr]
}

func main() {
	sl := &SpinLock{mem: []int{0}, ptr: 0, nv: 0}

	for i := 0; i < 2017; i++ {
		sl.Next(316)
	}

	fmt.Println(sl.GetValue(1))

}
