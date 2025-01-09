package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"strconv"
	"strings"
)

type Layers struct {
	maxLayerDepth int
	layers        map[int]*Scanner
}

func NewLayers() *Layers {
	l := &Layers{}
	l.layers = make(map[int]*Scanner)
	return l
}

func (l *Layers) GetMaxLayerDepth() int {
	return l.maxLayerDepth
}

func (l *Layers) AddSecurityScanner(s *Scanner) {
	if s.dep > l.maxLayerDepth {
		l.maxLayerDepth = s.dep
	}
	l.layers[s.dep] = s
}

func (l *Layers) HasScannerDepthAtPos(depth, pos int) (bool, int, int) {
	if s, e := l.layers[depth]; e {
		if s.Pos() == pos {
			return true, s.dep, s.ran
		}
	}
	return false, 0, 0
}

func (l *Layers) Advance() {
	for _, layer := range l.layers {
		layer.Advance()
	}
}

type Scanner struct {
	ptr int // index
	dep int // depth
	ran int // range
	dir int // direction 1 or -1
}

func (s *Scanner) Advance() {
	s.ptr += s.dir
	switch s.dir {
	case 1:
		if s.ptr == s.ran {
			s.dir = -1
			s.ptr += 2 * s.dir

		}
	case -1:
		if s.ptr == -1 {
			s.dir = 1
			s.ptr += 2 * s.dir
		}
	}
}

func (s *Scanner) Pos() int {
	return s.ptr
}

func main() {
	layers := getData("../data.txt")
	depth := -1
	maxDepth := layers.GetMaxLayerDepth()
	severity := 0
	for depth <= maxDepth {
		depth++
		if has, d, r := layers.HasScannerDepthAtPos(depth, 0); has {
			severity += d * r
		}
		layers.Advance()
	}

	fmt.Println(severity)
}

func getData(filename string) *Layers {
	lines, _ := file.GetLines(filename)
	layers := NewLayers()
	for _, line := range lines {
		tokens := strings.Fields(line)
		tokens = strings.Split(strings.Join(tokens, ""), ":")
		layers.AddSecurityScanner(&Scanner{dir: 1, dep: getIntVal(tokens[0]), ran: getIntVal(tokens[1])})
	}
	return layers
}

func getIntVal(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}
