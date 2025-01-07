package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"strings"
)

type List struct {
	data []byte
	skip int
	ptr  int
}

func (l *List) Process(input []byte) {
	for i := 0; i < len(input); i++ {
		if int(input[i]) <= len(l.data) { // input must be valid
			data := l.getData(int(input[i]))
			l.reverse(data)
			l.setData(data)
			l.advance(int(input[i]) + l.skip)
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

func (l *List) getData(length int) []byte {
	data := make([]byte, length)
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

func (l *List) setData(data []byte) {
	for i, j := l.ptr, 0; j < len(data); i, j = i+1, j+1 {
		if i >= len(l.data) {
			for i >= len(l.data) {
				i -= len(l.data)
			}
		}
		l.data[i] = data[j]
	}
}

func (l *List) reverse(data []byte) {
	for i, j, h := 0, len(data)-1, len(data)/2; i < h; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func (l *List) DenseHash() string {
	data := make([]byte, 0, len(l.data)/16)

	for i := 0; i < len(l.data); i += 16 {
		val := l.data[i]
		for j := i + 1; j < i+16; j++ {
			val ^= l.data[j]
		}
		data = append(data, val)
	}

	sb := strings.Builder{}
	for i := 0; i < len(data); i++ {
		sb.WriteString(fmt.Sprintf("%02x", data[i]))
	}

	return sb.String()
}

// not right: 9d5f4561367d379cfbf04f8c471c095
func main() {
	args := []struct {
		size  int
		input string
	}{
		{size: 256, input: "../test2.txt"},
		{size: 256, input: "../test3.txt"},
		{size: 256, input: "../test4.txt"},
		{size: 256, input: "../test5.txt"},
		{size: 256, input: "../data.txt"},
	}
	arg := args[4]

	list := genList(arg.size)
	input := getInputs(arg.input)

	fmt.Println(list, input)
	for i := 0; i < 64; i++ {
		list.Process(input)
	}

	fmt.Println(list, input)
	fmt.Println(list.DenseHash())

}

func genList(size int) *List {
	list := List{data: make([]byte, size)}
	for i := 0; i < size; i++ {
		list.data[i] = byte(i)
	}
	return &list
}

func getInputs(filename string) []byte {
	content, _ := file.GetContent(filename)
	suffix := []int{17, 31, 73, 47, 23}
	cs := strings.TrimSpace(string(content))
	input := make([]byte, 0, len(cs)+len(suffix))
	for i := range cs {
		input = append(input, byte(cs[i]))
	}
	for i := range suffix {
		input = append(input, byte(suffix[i]))
	}
	return input
}
