package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"slices"
)

const (
	space  = ' '
	vert   = '|'
	hor    = '-'
	corner = '+'
)

var (
	N        = Pos{-1, 0}
	E        = Pos{0, 1}
	S        = Pos{1, 0}
	W        = Pos{0, -1}
	dirs     = []Pos{S, W, N, E}
	opp      = map[Pos]Pos{N: S, S: N, W: E, E: W}
	gridChar = []byte{vert, hor, corner}
)

type Grid [][]byte

type Pos struct {
	Y int
	X int
}

func (p Pos) Add(o Pos) Pos {
	return Pos{Y: p.Y + o.Y, X: p.X + o.X}
}

func (g Grid) Has(pos Pos) bool {
	if pos.Y >= 0 && pos.Y < len(g) && pos.X >= 0 && pos.X < len(g[pos.Y]) {
		if g[pos.Y][pos.X] != space {
			return true
		}
	}
	return false
}

func (g Grid) GridChar(pos Pos) bool {
	if g.Has(pos) {
		if slices.Contains(gridChar, g[pos.Y][pos.X]) {
			return true
		}
	}
	return false
}

func (g Grid) Start() Pos {
	var pos Pos
	for x := 0; x < len(g[0]); x++ {
		if g[0][x] != space {
			pos.X = x
			break
		}
	}
	return pos
}

func (g Grid) Next(pos Pos, dir Pos) (*Pos, *Pos) {
	var nextPos, nextDir *Pos

	nextDirs := []Pos{dir}
	for _, d := range dirs {
		if d != dir && d != opp[dir] {
			nextDirs = append(nextDirs, d)
		}
	}

	for _, d := range nextDirs {
		nd := d
		np := pos.Add(d)
		if g.Has(np) {
			nextPos, nextDir = &np, &nd
			break
		}
	}

	return nextPos, nextDir
}

func main() {

	grid := getData("../data.txt")
	d := dirs[0]
	p := grid.Start()
	var letters []byte

	if !grid.GridChar(p) {
		letters = append(letters, grid[p.Y][p.X])
	}

	np, nd := grid.Next(p, d)
	for np != nil {
		p = *np
		d = *nd
		if !grid.GridChar(p) {
			letters = append(letters, grid[p.Y][p.X])
		}
		np, nd = grid.Next(p, d)
	}

	fmt.Println(string(letters))
}

func getData(filename string) Grid {
	lines, _ := file.GetLines(filename)
	grid := make(Grid, len(lines))
	for y, line := range lines {
		grid[y] = []byte(line)
	}
	return grid
}
