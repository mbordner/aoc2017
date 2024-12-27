package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common"
	"github.com/mbordner/aoc2017/common/file"
)

func main() {

	content, _ := file.GetContent("../data.txt")

	l := len(content)
	content = append(content, content...)
	sum := 0
	for i := 0; i < l; i++ {
		if content[i] == content[i+1] {
			sum += common.ByteCharToInt(content[i])
		}
	}

	fmt.Println(sum)
}
