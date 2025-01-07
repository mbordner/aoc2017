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
	parent   *Program
	name     string
	num      int
	children []string
}

func main() {
	root := getData("../data.txt")
	fmt.Println(root.name)
}

func getData(filename string) *Program {
	programs := make(map[string]*Program)

	lines, _ := file.GetLines(filename)
	for _, line := range lines {
		matches := reProgram.FindStringSubmatch(line)
		program := &Program{name: matches[1], parent: nil}
		val, _ := strconv.ParseInt(matches[2], 10, 64)
		program.num = int(val)
		if matches[3] != "" {
			program.children = strings.Split(matches[3], ", ")
		} else {
			program.children = []string{}
		}
		programs[program.name] = program
	}

	for _, program := range programs {
		for _, child := range program.children {
			programs[child].parent = program
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
