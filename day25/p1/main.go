package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"regexp"
	"strconv"
	"strings"
)

var (
	reState  = regexp.MustCompile(`In state (\w+):\n.\s+If the current value is 0:\n\s+- Write the value (1|0)\.\n\s+- Move one slot to the (left|right)\.\n\s+- Continue with state (\w+)\.\n\s+If the current value is 1:\n\s+- Write the value (1|0)\.\n\s+- Move one slot to the (left|right)\.\n\s+- Continue with state (\w+)\.`)
	reHeader = regexp.MustCompile(`Begin in state (\w+).\nPerform a diagnostic checksum after (\d+) steps\.`)
)

func main() {
	tm := getData("../data.txt")
	for i := 0; i < tm.steps; i++ {
		tm.Step()
	}
	fmt.Println(tm.CheckSum())
}

type TuringMachineActions struct {
	write bool
	dir   string
	next  string
}

type TuringMachineState struct {
	name string
	on   TuringMachineActions
	off  TuringMachineActions
}

type TuringMachine struct {
	cur    string
	steps  int
	values map[int]bool
	states map[string]TuringMachineState
	cursor int
}

func (tm *TuringMachine) Step() {
	curState := tm.states[tm.cur]

	var actions TuringMachineActions
	if tm.GetVal() {
		actions = curState.on
	} else {
		actions = curState.off
	}

	tm.SetVal(actions.write)

	if actions.dir == "left" {
		tm.cursor--
	} else {
		tm.cursor++
	}

	tm.cur = actions.next
}

func (tm *TuringMachine) GetVal() bool {
	if b, e := tm.values[tm.cursor]; e {
		return b
	}
	return false
}

func (tm *TuringMachine) SetVal(v bool) {
	tm.values[tm.cursor] = v
}

func (tm *TuringMachine) CheckSum() int {
	count := 0
	for _, b := range tm.values {
		if b {
			count++
		}
	}
	return count
}

func getData(filename string) *TuringMachine {
	content, _ := file.GetContent(filename)
	tokens := strings.Split(strings.TrimSpace(string(content)), "\n\n")

	tm := &TuringMachine{}
	matches := reHeader.FindStringSubmatch(tokens[0])
	tm.cur = matches[1]
	tm.steps = atoi(matches[2])
	tm.values = make(map[int]bool)
	tm.states = make(map[string]TuringMachineState)

	for _, str := range tokens[1:] {
		matches = reState.FindStringSubmatch(str)
		state := TuringMachineState{
			name: matches[1],
		}

		state.off = TuringMachineActions{}
		if matches[2] == "1" {
			state.off.write = true
		} else {
			state.off.write = false
		}
		state.off.dir = matches[3]
		state.off.next = matches[4]

		state.on = TuringMachineActions{}
		if matches[5] == "1" {
			state.on.write = true
		} else {
			state.on.write = false
		}
		state.on.dir = matches[6]
		state.on.next = matches[7]

		tm.states[state.name] = state
	}

	return tm
}

func atoi(s string) int {
	val, _ := strconv.ParseInt(s, 10, 64)
	return int(val)
}
