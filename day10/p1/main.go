package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"strconv"
	"strings"
)

type List struct {
	data []int
	skip int
	ptr  int
}

func main() {
	args := []struct {
		size  int
		input string
	}{
		{size: 5, input: "../test.txt"},
	}
	arg := args[0]

	list := genList(arg.size)
	input := getInputs(arg.input)

	fmt.Println(list, input)
}

func genList(size int) List {
	list := List{data: make([]int, size)}
	for i := 0; i < size; i++ {
		list.data[i] = i
	}
	return list
}

func getInputs(filename string) []int {
	content, _ := file.GetContent(filename)
	tokens := strings.Fields(string(content))
	cs := strings.Join(tokens, "")
	tokens = strings.Split(cs, ",")
	input := make([]int, len(tokens))
	for i := range input {
		input[i], _ = strconv.Atoi(tokens[i])
	}
	return input
}
