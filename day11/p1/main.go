package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common"
	"github.com/mbordner/aoc2017/common/file"
	"github.com/mbordner/aoc2017/common/hexagon"
)

func main() {
	dirs := getData("../data.txt")
	start := hexagon.Cell[int]{Q: 0, R: 0, S: 0}
	cur := start
	for _, d := range dirs {
		cur = cur.Next(d)
	}
	goal := cur

	queue := make(common.Queue[hexagon.Cell[int]], 100)
	visited := make(hexagon.Grid[int])
	prev := make(hexagon.CellLinker[int])

	visited[start] = true
	queue.Enqueue(start)

	var solution hexagon.Cells[int]

	for !queue.Empty() {
		cur = *(queue.Dequeue())
		if cur == goal {
			solution = hexagon.Cells[int]{}
			for p := cur; p != start; p = prev[p] {
				solution = append(hexagon.Cells[int]{p}, solution...)
			}
			break
		} else {
			ns := cur.Neighbors()
			for _, n := range ns {
				if !visited.Has(n) {
					visited[n] = true
					prev[n] = cur
					queue.Enqueue(n)
				}
			}
		}
	}

	fmt.Println(start, goal)
	fmt.Println(solution)
	fmt.Println(len(solution))
}

func getData(filename string) hexagon.Dirs {
	content, _ := file.GetContent(filename)
	return hexagon.DirsFromString(string(content))
}
