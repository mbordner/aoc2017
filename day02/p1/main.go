package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"sort"
	"strconv"
	"strings"
)

func main() {

	sum := 0
	lines, _ := file.GetLines("../data.txt")
	for _, line := range lines {
		tokens := strings.Fields(line)
		vals := make([]int, len(tokens))
		for i, t := range tokens {
			vals[i], _ = strconv.Atoi(t)
		}
		sort.Ints(vals)
		sum += vals[len(vals)-1] - vals[0]
	}

	fmt.Println(sum)
}
