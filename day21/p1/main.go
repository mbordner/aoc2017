package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"github.com/mbordner/aoc2017/common/grid"
	"strings"
)

type Enhancements map[string]string

func main() {
	enhancements := getEnhancements("../data.txt")

	start := grid.ExpandGrid([]byte(`.#./..#/###`), '/')
	cur := start.Clone()

	for i := 0; i < 5; i++ {
		cur = enhance(cur, enhancements)
	}

	fmt.Println(on(cur))
}

func on(g grid.Grid[byte]) int {
	count := 0
	for y := range g {
		for x := range g[y] {
			if g[y][x] == '#' {
				count++
			}
		}
	}
	return count
}

func gsz(g grid.Grid[byte]) int {
	if len(g)%2 == 0 {
		return 2
	} else if len(g)%3 == 0 {
		return 3
	}
	panic("invalid grid size")
}

func enhance(g grid.Grid[byte], enhancements Enhancements) grid.Grid[byte] {
	var ng grid.Grid[byte]
	var ngsz int
	cgsz := gsz(g)
	switch cgsz {
	case 2:
		ngsz = len(g) / 2 * 3
		ng = grid.NewGrid[byte](ngsz, ngsz, '.')
	case 3:
		ngsz = len(g) / 3 * 4
		ng = grid.NewGrid[byte](ngsz, ngsz, '.')
	default:
		panic("invalid grid size")
	}
	for y, ny := 0, 0; y < len(g); y, ny = y+cgsz, ny+cgsz+1 {
		for x, nx := 0, 0; x < len(g[y]); x, nx = x+cgsz, nx+cgsz+1 {
			tg := g.Read(x, y, x+cgsz-1, y+cgsz-1)
			tgc := tg.Condense('/')
			ntg := grid.ExpandGrid([]byte(enhancements[string(tgc)]), '/')
			ng.Write(nx, ny, ntg)
		}
	}
	return ng
}

func getEnhancements(filename string) Enhancements {
	enhancements := make(Enhancements)
	lines, _ := file.GetLines(filename)
	for _, line := range lines {
		tokens := strings.Split(strings.Join(strings.Fields(line), ""), "=>")
		enhancement := tokens[1]
		variations := getGridMatchVariations(tokens[0])
		for _, variation := range variations {
			if en, ex := enhancements[variation]; ex {
				if en != enhancement {
					panic("enhancements mismatch")
				}
			} else {
				enhancements[variation] = enhancement
			}
		}
	}
	return enhancements
}

func getGridMatchVariations(s string) []string {
	keys := make(map[string]bool)
	g := grid.ExpandGrid([]byte(s), '/')
	grids := []grid.Grid[byte]{g, g.FlipHorizontal(), g.FlipVertical()}
	for _, o := range grids {
		keys[string(o.Condense('/'))] = true
		for i := 0; i < 3; i++ {
			o = o.RotateRight()
			keys[string(o.Condense('/'))] = true
		}
	}
	vals := make([]string, 0, len(keys))
	for key := range keys {
		vals = append(vals, key)
	}
	return vals
}
