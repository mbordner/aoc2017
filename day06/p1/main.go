package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"strconv"
	"strings"
)

func main() {
	mb := getData("../data.txt")
	allocations := make(map[string]bool)
	allocations[mb.String()] = true

	count := 0
	for {
		mb.Reallocate()
		count++
		allocation := mb.String()
		if _, e := allocations[allocation]; e {
			break
		}
		allocations[allocation] = true
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
