package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"strings"
)

func main() {
	count := 0
	lines, _ := file.GetLines("../data.txt")
	for _, line := range lines {
		if checkPassPhrase(line) {
			count++
		}
	}

	fmt.Println(count)
}

func checkPassPhrase(passPhrase string) bool {
	words := strings.Fields(passPhrase)
	wordsMap := make(map[string]bool)

	for _, word := range words {
		if _, e := wordsMap[word]; e {
			return false
		}
		wordsMap[word] = true
	}

	return true
}
