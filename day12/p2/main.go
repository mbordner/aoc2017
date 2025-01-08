package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common"
	"github.com/mbordner/aoc2017/common/file"
	"github.com/mbordner/aoc2017/common/graph"
	"strconv"
	"strings"
)

type NodeMap map[int]*graph.Node

func main() {
	g := getData("../data.txt")

	seen := make(NodeMap)
	var groups []NodeMap

	nodes := g.GetNodes()

	for _, n := range nodes {
		if _, e := seen[n.GetID().(int)]; !e {
			queue := make(common.Queue[*graph.Node], 0, g.Len())
			group := make(NodeMap)

			seen[n.GetID().(int)] = n
			group[n.GetID().(int)] = n
			queue.Enqueue(n)

			for !queue.Empty() {
				cur := *(queue.Dequeue())
				edges := cur.GetEdges()
				for _, edge := range edges {
					o := edge.GetDestination()
					if _, e = seen[o.GetID().(int)]; !e {
						seen[o.GetID().(int)] = o
						group[o.GetID().(int)] = o
						queue.Enqueue(o)
					}
				}
			}

			groups = append(groups, group)
		}

	}

	fmt.Println(len(groups))
}

func getData(filename string) *graph.Graph {
	g := graph.NewGraph()

	lines, _ := file.GetLines(filename)
	for _, line := range lines {
		tokens := strings.Split(line, " <-> ")
		nId := getIntVal(tokens[0])
		tokens = strings.Split(strings.Join(strings.Fields(tokens[1]), ""), ",")
		var n *graph.Node
		if n = g.GetNode(nId); n == nil {
			n = g.CreateNode(nId)
		}

		for _, t := range tokens {
			oId := getIntVal(t)
			var o *graph.Node
			if o = g.GetNode(oId); o == nil {
				o = g.CreateNode(oId)
			}
			n.AddEdge(o, 1)
		}
	}

	return g
}

func getIntVal(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}
