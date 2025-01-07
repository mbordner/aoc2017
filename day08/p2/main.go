package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"regexp"
	"strconv"
)

var (
	reStatement = regexp.MustCompile(`(\w+)\s+(inc|dec)\s+(-?\d+)\s+if\s+(\w+)\s+(>|<|>=|<=|==|!=)\s+(-?\d+)`)
	reDigits    = regexp.MustCompile(`(-?\d+)`)
)

type Registers struct {
	maxReg string
	maxVal int
	regs   map[string]int
}

func (r *Registers) getVal(s string) int {
	if reDigits.MatchString(s) {
		return r.getIntVal(s)
	}
	return r.getRegVal(s)
}

func (r *Registers) getRegVal(reg string) int {
	if val, e := r.regs[reg]; e {
		return val
	}
	return 0
}

func (r *Registers) setRegVal(reg string, val int) {
	r.regs[reg] = val
}

func (r *Registers) GetCurrentLargestRegValue() (string, int) {
	maxVal := 0
	maxReg := ""
	for r, v := range r.regs {
		if v > maxVal {
			maxVal = v
			maxReg = r
		}
	}
	return maxReg, maxVal
}

func (r *Registers) getIntVal(s string) int {
	val, _ := strconv.ParseInt(s, 10, 64)
	return int(val)
}

func (r *Registers) condition(left, op, right string) bool {
	leftVal, rightVal := r.getVal(left), r.getVal(right)
	switch op {
	case ">":
		return leftVal > rightVal
	case "<":
		return leftVal < rightVal
	case "==":
		return leftVal == rightVal
	case ">=":
		return leftVal >= rightVal
	case "<=":
		return leftVal <= rightVal
	case "!=":
		return leftVal != rightVal
	}
	panic("unknown operation")
}

func (r *Registers) ProcessStatement(stmt string) bool {
	if reStatement.MatchString(stmt) {
		matches := reStatement.FindStringSubmatch(stmt)
		if r.condition(matches[4], matches[5], matches[6]) {
			rVal := r.getVal(matches[3])
			regVal := r.getRegVal(matches[1])
			switch matches[2] {
			case "inc":
				r.setRegVal(matches[1], regVal+rVal)
			case "dec":
				r.setRegVal(matches[1], regVal-rVal)
			}
			mr, mrv := r.GetCurrentLargestRegValue()
			if mrv > r.maxVal {
				r.maxVal = mrv
				r.maxReg = mr
			}
		}
		return true
	}
	return false
}

func (r *Registers) GetMaxRegValueDuringProcessing() (string, int) {
	return r.maxReg, r.maxVal
}

func main() {
	regs := &Registers{regs: make(map[string]int)}
	lines, _ := file.GetLines("../data.txt")
	for _, line := range lines {
		regs.ProcessStatement(line)
	}
	fmt.Println(regs.GetMaxRegValueDuringProcessing())
}
