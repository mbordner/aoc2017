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

func reverseTurnDir(d Pos) Pos {
	return getTurnDir(getTurnDir(d, R), R)
}

type Pos struct {
	Y int
	X int
}

func (p Pos) Add(o Pos) Pos {
	return Pos{Y: p.Y + o.Y, X: p.X + o.X}
}

type VirusCarrier struct {
	p   Pos
	d   Pos
	bic int
}

type GridState int

const (
	CLEAN GridState = iota
	INFECTED
	WEAKENED
	FLAGGED
)

type Grid map[Pos]GridState

func (g Grid) SetState(p Pos, s GridState) {
	g[p] = s
}

func (g Grid) GetState(p Pos) GridState {
	if s, e := g[p]; e {
		return s
	}
	return CLEAN
}

func (vc *VirusCarrier) Burst(g Grid) {
	curState := g.GetState(vc.p)
	var nextState GridState
	switch curState {
	case CLEAN:
		vc.d = getTurnDir(vc.d, L)
		nextState = WEAKENED
	case INFECTED:
		vc.d = getTurnDir(vc.d, R)
		nextState = FLAGGED
	case FLAGGED:
		vc.d = reverseTurnDir(vc.d)
		nextState = CLEAN
	case WEAKENED:
		nextState = INFECTED
	}

	g.SetState(vc.p, nextState)
	if nextState == INFECTED {
		vc.bic++
	}
	vc.p = vc.p.Add(vc.d)
}

func main() {
	grid, vc := getData("../data.txt")

	for i := 0; i < 10000000; i++ {
		vc.Burst(grid)
	}

	fmt.Println(vc.bic)
}

func getData(filename string) (Grid, *VirusCarrier) {
	lines, _ := file.GetLines(filename)
	grid := make(Grid)
	for y, line := range lines {
		for x, b := range []byte(line) {
			if b == '#' {
				grid.SetState(Pos{Y: y, X: x}, INFECTED)
			}
		}
	}
	vcPos := Pos{Y: len(lines) / 2, X: len(lines) / 2}
	vc := VirusCarrier{p: vcPos, d: N}
	return grid, &vc
}
