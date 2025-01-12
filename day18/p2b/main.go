package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common"
	"github.com/mbordner/aoc2017/common/file"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	reDigits = regexp.MustCompile(`(-?\d+)`)
)

type State int

const (
	RUNNING State = iota
	RCV
	SND
	EXITED
)

type Computer struct {
	regs      map[string]int
	ptr       int
	program   []string
	id        int
	state     State
	evaluated common.Stack[string]
}

func (c *Computer) Load(filename string) *Computer {
	c.ptr = 0
	c.regs = make(map[string]int)
	c.program, _ = file.GetLines(filename)
	c.evaluated = make(common.Stack[string], 0, 10)
	return c
}

func (c *Computer) SetID(id int) {
	c.id = id
	c.regs["p"] = id
}

// Run when val is non nil, it's receiving a value, when it returns nil, it's waiting for receive
// when returns non nil it's sending a value
// bool return true when run is completed naturally
func (c *Computer) Run(val *int) (*int, bool) {
	c.state = RUNNING
	for c.ptr >= 0 && c.ptr < len(c.program) {
		instr := c.program[c.ptr]
		c.evaluated.Push(instr)
		c.ptr++
		tokens := strings.Fields(instr)
		switch tokens[0] {
		case "snd":
			v := c.GetValue(tokens[1])
			c.state = SND
			return &v, false // we want to advance pointer on send
		case "set":
			c.SetValue(tokens[1], c.GetValue(tokens[2]))
		case "add":
			c.SetValue(tokens[1], c.GetValue(tokens[1])+c.GetValue(tokens[2]))
		case "mul":
			c.SetValue(tokens[1], c.GetValue(tokens[1])*c.GetValue(tokens[2]))
		case "mod":
			c.SetValue(tokens[1], c.GetValue(tokens[1])%c.GetValue(tokens[2]))
		case "rcv":
			if val != nil {
				v := *val
				c.SetValue(tokens[1], v)
				val = nil // clear val after this so next rcv breaks
			} else {
				c.ptr-- // set ptr back to this rcv
				c.state = RCV
				c.evaluated.Pop()
				return nil, false
			}
		case "jgz":
			if c.GetValue(tokens[1]) > 0 {
				c.ptr--
				c.ptr += c.GetValue(tokens[2])
			}
		}
	}
	c.state = EXITED
	return nil, true
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
	c0, c1 := &Computer{}, &Computer{}
	filename := "../data.txt"
	c0.Load(filename)
	c1.Load(filename)
	c0.SetID(0)
	c1.SetID(1)

	q0 := make(common.Queue[int], 0, 10)
	q1 := make(common.Queue[int], 0, 10)

	canContinue := func() bool {
		if slices.Contains([]State{RUNNING, SND}, c0.state) {
			return true
		}
		if slices.Contains([]State{RUNNING, SND}, c1.state) {
			return true
		}
		if c0.state == RCV && !q1.Empty() {
			return true
		}
		if c1.state == RCV && !q0.Empty() {
			return true
		}
		return false
	}

	for canContinue() {
		var exited bool
		var val *int
		if c0.state == RCV && !q1.Empty() {
			val = q1.Dequeue()
			val, exited = c0.Run(val)
			if !exited {
				if val != nil {
					q0.Enqueue(*val)
				}
			}
		}
		if c1.state == RCV && !q0.Empty() {
			val = q0.Dequeue()
			val, exited = c1.Run(val)
			if !exited {
				if val != nil {
					q1.Enqueue(*val)
				}
			}
		}
		if slices.Contains([]State{RUNNING, SND}, c0.state) {
			val, exited = c0.Run(nil)
			if !exited {
				if val != nil {
					q0.Enqueue(*val)
				}
			}
		}
		if slices.Contains([]State{RUNNING, SND}, c1.state) {
			val, exited = c1.Run(nil)
			if !exited {
				if val != nil {
					q1.Enqueue(*val)
				}
			}
		}
	}

	count := 0
	for _, instr := range c1.evaluated {
		if strings.HasPrefix(instr, "snd") {
			count++
		}
	}

	fmt.Println(count)

}

func atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}
