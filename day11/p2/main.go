package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"github.com/mbordner/aoc2017/common/hexagon"
)

func main() {
	dirs := getData("../data.txt")

	start := hexagon.Cell[int]{Q: 0, R: 0, S: 0}

	visited := make(hexagon.Grid[int])
	visited[start] = true

	maxDistance := 0
	var maxCell hexagon.Cell[int]

	cur := start
	for _, d := range dirs {
		cur = cur.Next(d)
		if !visited.Has(cur) {
			visited[cur] = true
			distance := start.Distance(cur)
			if distance > maxDistance {
				maxDistance = distance
				maxCell = cur
			}
		}
	}

	goal := cur

	fmt.Printf("start: %v\n", start)
	fmt.Printf("goal: %v\n", goal)
	fmt.Printf("max distance cell: %v\n", maxCell)
	fmt.Printf("number of steps: %d\n", len(dirs))
	fmt.Printf("distance to goal was: %d\n", start.Distance(goal))
	fmt.Printf("visited %d cells\n", len(visited))
	fmt.Printf("max distance %d\n", maxDistance)

}

func getData(filename string) hexagon.Dirs {
	content, _ := file.GetContent(filename)
	return hexagon.DirsFromString(string(content))
}
