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
	regs    map[string]int64
	ptr     int
	program []string
}

func (c *Computer) Load(filename string) *Computer {
	c.ptr = 0
	c.regs = make(map[string]int64)
	c.program, _ = file.GetLines(filename)
	return c
}

func (c *Computer) Run() {
	for c.ptr >= 0 && c.ptr < len(c.program) {
		instr := c.program[c.ptr]
		fmt.Printf("%d: %s %v\n", c.ptr+1, instr, c.regs)
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
				c.ptr += int(c.GetValue(tokens[2]))
			}
		case "jnz":
			if c.GetValue(tokens[1]) != 0 {
				c.ptr--
				c.ptr += int(c.GetValue(tokens[2]))
			}
		}
	}
}

func (c *Computer) GetValue(s string) int64 {
	if reDigits.MatchString(s) {
		matches := reDigits.FindStringSubmatch(s)
		return atoi(matches[1])
	}
	if v, e := c.regs[s]; e {
		return v
	}
	return 0
}

func (c *Computer) SetValue(r string, v int64) {
	if !reDigits.MatchString(r) {
		c.regs[r] = v
	}
}

func main() {
	c := &Computer{}
	c.Load("../data.txt")
	c.SetValue("a", 1)
	c.Run()
}

func atoi(s string) int64 {
	val, _ := strconv.ParseInt(s, 10, 64)
	return val
}
