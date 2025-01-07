package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"regexp"
	"strconv"
	"strings"
)

var (
	reProgram = regexp.MustCompile(`(\w+)\s+\((\d+)\)(?:\s+->\s+(.*))?`)
)

type Program struct {
	parent     *Program
	name       string
	weight     int
	childNames []string
	children   []*Program
}

func (p *Program) FindUnbalanced() (*Program, int) {
	if len(p.children) == 0 {
		return nil, 0
	}
	sum := 0
	cws := make([]int, len(p.children))
	valCounts := make(map[int]int)
	for i, c := range p.children {
		cw, gcw := c.Weight()
		cws[i] = cw + gcw
		sum += cws[i]
		if vc, e := valCounts[cws[i]]; e {
			valCounts[cws[i]] = vc + 1
		} else {
			valCounts[cws[i]] = 1
		}
	}

	if len(valCounts) > 1 {
		var balancedWeight int
		for w, c := range valCounts {
			if c != 1 {
				balancedWeight = w
				break
			}
		}
		for i, c := range p.children {
			if balancedWeight != cws[i] {
				ubp, delta := c.FindUnbalanced()
				if ubp != nil {
					return ubp, delta
				}
				cw, gcw := c.Weight()
				return c, cw - (balancedWeight - gcw)
			}
		}
	}

	return nil, 0
}

func (p *Program) Weight() (int, int) {
	childrenWeight := 0
	for _, child := range p.children {
		childWeight, grandChildrenWeight := child.Weight()
		childrenWeight += childWeight + grandChildrenWeight
	}
	return p.weight, childrenWeight
}

func main() {
	root := getData("../data.txt")

	unbalanced, delta := root.FindUnbalanced()
	ubw, _ := unbalanced.Weight()
	fmt.Println(ubw - delta)
}

func getData(filename string) *Program {
	programs := make(map[string]*Program)

	lines, _ := file.GetLines(filename)
	for _, line := range lines {
		matches := reProgram.FindStringSubmatch(line)
		program := &Program{name: matches[1], parent: nil}
		val, _ := strconv.ParseInt(matches[2], 10, 64)
		program.weight = int(val)
		if matches[3] != "" {
			program.childNames = strings.Split(matches[3], ", ")
		} else {
			program.childNames = []string{}
		}
		programs[program.name] = program
	}

	for _, program := range programs {
		program.children = make([]*Program, 0, len(program.childNames))
		for _, childName := range program.childNames {
			programs[childName].parent = program
			program.children = append(program.children, programs[childName])
		}
	}

	var root *Program
	for _, program := range programs {
		if program.parent == nil {
			if root != nil {
				fmt.Println("more than one root program")
			}
			root = program
		}
	}

	return root
}
