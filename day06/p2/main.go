package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"strconv"
	"strings"
)

func main() {
	mb := getData("../data.txt")
	allocations := make(map[string]int)

	counts := make(map[string]int)
	count := 0
	for {
		mb.Reallocate()
		for a, c := range counts {
			counts[a] = c + 1
		}
		allocation := mb.String()
		if v, e := allocations[allocation]; e {
			allocations[allocation] = v + 1
		} else {
			allocations[allocation] = 1
		}
		if allocations[allocation] == 2 {
			counts[allocation] = 0
		} else if allocations[allocation] == 3 {
			count = counts[allocation]
			break
		}
	}

	fmt.Println(count)
}

type MemoryBank struct {
	blocks []int
}

func (mb *MemoryBank) String() string {
	blocks := make([]string, len(mb.blocks))
	for i, b := range mb.blocks {
		blocks[i] = fmt.Sprintf("%d", b)
	}
	return strings.Join(blocks, ",")
}

func (mb *MemoryBank) Reallocate() {
	val, index := mb.blocks[0], 0
	for i, b := range mb.blocks[1:] {
		if b > val {
			val, index = b, i+1
		}
	}

	mb.blocks[index] = 0
	for val > 0 {
		index++
		if index == len(mb.blocks) {
			index = 0
		}
		mb.blocks[index]++
		val--
	}
}

func getData(filename string) *MemoryBank {
	content, _ := file.GetContent(filename)
	tokens := strings.Fields(string(content))
	mb := MemoryBank{}
	mb.blocks = make([]int, len(tokens))
	for i, token := range tokens {
		val, _ := strconv.ParseInt(token, 10, 64)
		mb.blocks[i] = int(val)
	}
	return &mb
}
