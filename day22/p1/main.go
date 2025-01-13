package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
)

var (
	N = Pos{-1, 0}
	E = Pos{0, 1}
	S = Pos{1, 0}
	W = Pos{0, -1}
)

type Turn int

const (
	UNKNOWN Turn = iota
	L
	R
)

var turns = map[Pos]map[Turn]Pos{
	N: {L: W, R: E},
	E: {L: N, R: S},
	S: {L: E, R: W},
	W: {L: S, R: N},
}

func getTurnDir(d Pos, t Turn) Pos {
	return turns[d][t]
}

type Pos struct {
	Y int
	X int
}

func (p Pos) Add(o Pos) Pos {
	return Pos{Y: p.Y + o.Y, X: p.X + o.X}
}

type VirusCarrier struct {
	p  Pos
	d  Pos
	bc int
}

type Grid map[Pos]bool

func (g Grid) IsInfected(p Pos) bool {
	if i, e := g[p]; e {
		return i
	}
	return false
}

func (g Grid) SetInfected(p Pos, i bool) {
	g[p] = i
}

func (vc *VirusCarrier) Burst(g Grid) {
	curInfected := g.IsInfected(vc.p)
	if curInfected {
		vc.d = getTurnDir(vc.d, R)
	} else {
		vc.d = getTurnDir(vc.d, L)
	}
	nextInfected := !curInfected
	g.SetInfected(vc.p, nextInfected)
	if nextInfected {
		vc.bc++
	}
	vc.p = vc.p.Add(vc.d)
}

func main() {
	grid, vc := getData("../data.txt")

	for i := 0; i < 10000; i++ {
		vc.Burst(grid)
	}

	fmt.Println(vc.bc)
}

func getData(filename string) (Grid, *VirusCarrier) {
	lines, _ := file.GetLines(filename)
	grid := make(Grid)
	for y, line := range lines {
		for x, b := range []byte(line) {
			if b == '#' {
				grid.SetInfected(Pos{Y: y, X: x}, true)
			}
		}
	}
	vcPos := Pos{Y: len(lines) / 2, X: len(lines) / 2}
	vc := VirusCarrier{p: vcPos, d: N}
	return grid, &vc
}
