package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	reDigits = regexp.MustCompile(`(-?\d+)`)
)

type Computer struct {
	regs    map[string]int
	ptr     int
	program []string
	played  []int
	Input   chan int
	Output  chan int
	id      int
	sent    int
}

func (c *Computer) Load(filename string) *Computer {
	c.ptr = 0
	c.regs = make(map[string]int)
	c.program, _ = file.GetLines(filename)
	c.played = []int{}
	c.Input = make(chan int, 80)
	c.Output = make(chan int, 80)
	return c
}

func (c *Computer) Run() {
	c.id = c.GetValue("p")
	c.sent = 0
	for c.ptr >= 0 && c.ptr < len(c.program) {
		instr := c.program[c.ptr]
		c.ptr++

		tokens := strings.Fields(instr)
		switch tokens[0] {
		case "snd":
			val := c.GetValue(tokens[1])
			c.sent++
			if c.id == 1 {
				fmt.Printf("%d (%d) sending %d\n", c.id, c.sent, val)
			}
			c.Output <- val
		case "set":
			c.SetValue(tokens[1], c.GetValue(tokens[2]))
		case "add":
			c.SetValue(tokens[1], c.GetValue(tokens[1])+c.GetValue(tokens[2]))
		case "mul":
			c.SetValue(tokens[1], c.GetValue(tokens[1])*c.GetValue(tokens[2]))
		case "mod":
			c.SetValue(tokens[1], c.GetValue(tokens[1])%c.GetValue(tokens[2]))
		case "rcv":
			//fmt.Printf("%d receiving\n", c.id)
			val := <-c.Input
			//fmt.Printf("%d received %d\n", c.id, val)
			c.SetValue(tokens[1], val)
		case "jgz":
			if c.GetValue(tokens[1]) > 0 {
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
	c0, c1 := &Computer{}, &Computer{}
	filename := "../data.txt"
	c0.Load(filename)
	c1.Load(filename)
	c0.SetValue("p", 0)
	c1.SetValue("p", 1)

	var wg sync.WaitGroup
	wg.Add(2)

	kill := make(chan bool)

	go func() {
	outer:
		for {
			select {
			case <-kill:
				break outer
			case c0Val := <-c0.Output:
				c1.Input <- c0Val
			case c1Val := <-c1.Output:
				c0.Input <- c1Val
			}
		}
	}()

	run := func(wg *sync.WaitGroup, c *Computer) {
		c.Run()
		wg.Done()
	}

	go run(&wg, c0)
	go run(&wg, c1)

	wg.Wait()
}

func atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}
