package main

import (
	"fmt"
	"math"
	"strings"
)

type Pos struct {
	Y int
	X int
}

func (p Pos) Add(o Pos) Pos {
	return Pos{Y: p.Y + o.Y, X: p.X + o.X}
}

type PosVal map[Pos]int

func (pv PosVal) Val(p Pos) int {
	if v, e := pv[p]; e {
		return v
	}
	return 0
}

func (pv PosVal) Print() {
	miX, miY, maX, maY := math.MaxUint32, math.MaxUint32, 0, 0
	maV := 0
	for p, v := range pv {
		if v > maV {
			maV = v
		}
		if p.X < miX {
			miX = p.X
		}
		if p.X > maX {
			maX = p.X
		}
		if p.Y < miY {
			miY = p.Y
		}
		if p.Y > maY {
			maY = p.Y
		}
	}
	rY, rX := maY-miY+1, maX-miX+1

	maxVStr := fmt.Sprintf("%d", maV)
	fmtStr := fmt.Sprintf("%%0%dd", len(maxVStr))

	grid := make([][]string, rY)
	for y := 0; y < rY; y++ {
		grid[y] = make([]string, rX)
		pY := miY + y
		for pX, x := miX, 0; pX <= maX; pX, x = pX+1, x+1 {
			p := Pos{X: pX, Y: pY}
			v := pv[p]
			grid[y][x] = fmt.Sprintf(fmtStr, v)
		}
	}

	for y := range grid {
		fmt.Println(strings.Join(grid[y], ","))
	}

}

func main() {
	deltas := []Pos{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}

	puzzleInput := 277678

	pv := make(PosVal)
	start := Pos{0, 0}
	pv[start] = 1

	level, squares, minVal, maxVal := 0, 1, 1, 1
outer:
	for {
		pv.Print()
		fmt.Println("-----")
		level += 1
		squares = (2*(level-1)+1)*4 + 4
		minVal = maxVal + 1
		maxVal = minVal + squares - 1
		levelEndPos := Pos{level, level}
		p := levelEndPos
		for i := minVal; i <= maxVal; i++ {
			leg := (i - minVal) / (squares / 4)
			var d Pos
			switch leg {
			case 0:
				d = Pos{Y: -1, X: 0}
			case 1:
				d = Pos{Y: 0, X: -1}
			case 2:
				d = Pos{Y: 1, X: 0}
			case 3:
				d = Pos{Y: 0, X: 1}
			}
			np := p.Add(d)

			sum := 0
			for _, delta := range deltas {
				sum += pv.Val(np.Add(delta))
			}

			pv[np] = sum
			if sum >= puzzleInput {
				pv.Print()
				fmt.Println("-----")
				fmt.Println(sum)
				break outer
			}

			p = np
		}
	}

}
