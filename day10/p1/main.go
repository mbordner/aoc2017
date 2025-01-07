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

func (l *List) Process(input []int) {
	for i := 0; i < len(input); i++ {
		if input[i] <= len(l.data) { // input must be valid
			data := l.getData(input[i])
			l.reverse(data)
			l.setData(data)
			l.advance(input[i] + l.skip)
			l.skip++
		}
	}
}

func (l *List) advance(amt int) {
	ptr := l.ptr + amt
	for ptr >= len(l.data) {
		ptr -= len(l.data)
	}
	l.ptr = ptr
}

func (l *List) getData(length int) []int {
	data := make([]int, length)
	for i, j := l.ptr, 0; j < length; i, j = i+1, j+1 {
		if i >= len(l.data) {
			for i >= len(l.data) {
				i -= len(l.data)
			}
		}
		data[j] = l.data[i]
	}
	return data
}

func (l *List) setData(data []int) {
	for i, j := l.ptr, 0; j < len(data); i, j = i+1, j+1 {
		if i >= len(l.data) {
			for i >= len(l.data) {
				i -= len(l.data)
			}
		}
		l.data[i] = data[j]
	}
}

func (l *List) reverse(data []int) {
	for i, j, h := 0, len(data)-1, len(data)/2; i < h; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func main() {
	args := []struct {
		size  int
		input string
	}{
		{size: 5, input: "../test.txt"},
		{size: 256, input: "../data.txt"},
	}
	arg := args[1]

	list := genList(arg.size)
	input := getInputs(arg.input)

	fmt.Println(list, input)
	list.Process(input)
	fmt.Println(list, input)
	fmt.Println(list.data[0] * list.data[1])
}

func genList(size int) *List {
	list := List{data: make([]int, size)}
	for i := 0; i < size; i++ {
		list.data[i] = i
	}
	return &list
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
