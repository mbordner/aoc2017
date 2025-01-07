package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
)

type Garbage struct {
	count int
	data  []byte
}

type Group struct {
	children []interface{}
}

func (g *Group) Add(c interface{}) {
	g.children = append(g.children, c)
}

func (g *Group) Score() int {
	return g.depthScore(1)
}

func (g *Group) depthScore(depth int) int {
	score := depth
	for _, child := range g.children {
		if group, ok := child.(*Group); ok {
			score += group.depthScore(depth + 1)
		}
	}
	return score
}

func (g *Group) GarbageLen() int {
	sum := 0
	for _, child := range g.children {
		if group, ok := child.(*Group); ok {
			sum += group.GarbageLen()
		} else {
			garbage := child.(Garbage)
			sum += garbage.count
		}
	}
	return sum
}

func NewGroup() *Group {
	g := &Group{}
	g.children = []interface{}{}
	return g
}

type GroupStack []*Group

func (gs *GroupStack) Push(group *Group) {
	*gs = append(*gs, group)
}

func (gs *GroupStack) Pop() *Group {
	g := gs.Peek()
	*gs = (*gs)[0 : len(*gs)-1]
	return g
}

func (gs *GroupStack) Peek() *Group {
	return (*gs)[len(*gs)-1]
}

func (gs *GroupStack) Empty() bool {
	return gs.Len() == 0
}

func (gs *GroupStack) Len() int {
	return len(*gs)
}

func main() {
	data, _ := file.GetContent("../data.txt")
	root := getStreamRootGroup(data)
	fmt.Println(root.GarbageLen())
}

func getStreamRootGroup(data []byte) *Group {
	gs := make(GroupStack, 0, 10)

	ptr := 0
	for ptr < len(data) {
		if data[ptr] == ',' { // still in cur group, going onto next child
			ptr++
		} else if data[ptr] == '{' {
			gs.Push(NewGroup())
			ptr++
		} else if data[ptr] == '}' {
			if gs.Len() > 1 {
				g := gs.Pop()
				gs.Peek().Add(g)
			}
			ptr++
		} else if data[ptr] == '<' {
			garbageStart := ptr
			gc := 0
			ptr++
			for ptr < len(data) {
				if data[ptr] == '!' {
					ptr += 2
				} else if data[ptr] == '>' {
					garbage := Garbage{count: gc, data: make([]byte, ptr+1-garbageStart)}
					copy(garbage.data, data[garbageStart:ptr+1])
					gs.Peek().Add(garbage)
					ptr++
					break
				} else {
					ptr++
					gc++
				}
			}
		}
	}

	return gs.Pop()
}

func getStreamGroupScore(data []byte) int {
	root := getStreamRootGroup(data)
	return root.Score()
}
