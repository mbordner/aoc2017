package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"regexp"
	"strconv"
	"strings"
)

var (
	reDigits = regexp.MustCompile(`(-?\d+)`)
)

type Computer struct {
	regs      map[string]int
	ptr       int
	program   []string
	evaluated []string
}

func (c *Computer) Load(filename string) *Computer {
	c.ptr = 0
	c.regs = make(map[string]int)
	c.program, _ = file.GetLines(filename)
	c.evaluated = []string{}
	return c
}

func (c *Computer) Run() {
	for c.ptr >= 0 && c.ptr < len(c.program) {
		instr := c.program[c.ptr]
		c.evaluated = append(c.evaluated, instr)
		c.ptr++
		tokens := strings.Fields(instr)
		switch tokens[0] {
		case "set":
			c.SetValue(tokens[1], c.GetValue(tokens[2]))
		case "sub":
			c.SetValue(tokens[1], c.GetValue(tokens[1])-c.GetValue(tokens[2]))
		case "add":
			c.SetValue(tokens[1], c.GetValue(tokens[1])+c.GetValue(tokens[2]))
		case "mul":
			c.SetValue(tokens[1], c.GetValue(tokens[1])*c.GetValue(tokens[2]))
		case "mod":
			c.SetValue(tokens[1], c.GetValue(tokens[1])%c.GetValue(tokens[2]))
		case "jgz":
			if c.GetValue(tokens[1]) > 0 {
				c.ptr--
				c.ptr += c.GetValue(tokens[2])
			}
		case "jnz":
			if c.GetValue(tokens[1]) != 0 {
				c.ptr--
				c.ptr += c.GetValue(tokens[2])
			}
		}
	}
}

func (c *Computer) GetValue(s string) int {
	if reDigits.MatchString(s) {
		matches := reDigits.FindStringSubmatch(s)
		return atoi(matches[1])
	}
	if v, e := c.regs[s]; e {
		return v
	}
	return 0
}

func (c *Computer) SetValue(r string, v int) {
	if !reDigits.MatchString(r) {
		c.regs[r] = v
	}
}

func main() {
	c := &Computer{}
	c.Load("../data.txt").Run()

	count := 0
	for _, instr := range c.evaluated {
		if strings.HasPrefix(instr, "mul") {
			count++
		}
	}
	fmt.Println(count)
}

func atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}
