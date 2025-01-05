package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"strconv"
)

type Jumps struct {
	vals []int
}

func (j *Jumps) Count() int {
	vals := make([]int, len(j.vals))
	copy(vals, j.vals)

	count := 0
	ptr := 0
	for ptr < len(vals) {
		last := ptr
		ptr += vals[ptr]
		count++
		if vals[last] >= 3 {
			vals[last]--
		} else {
			vals[last]++
		}
	}

	return count
}

func main() {
	jumps := getData("../data.txt")

	fmt.Println(jumps.Count())
}

func getData(filename string) Jumps {
	lines, _ := file.GetLines(filename)
	jumps := Jumps{}
	jumps.vals = make([]int, len(lines))
	for i, line := range lines {
		val, _ := strconv.Atoi(line)
		jumps.vals[i] = val
	}
	return jumps
}
