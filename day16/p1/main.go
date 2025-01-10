package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"regexp"
	"strconv"
	"strings"
)

var (
	reSpin     = regexp.MustCompile(`s(\d+)`)
	reExchange = regexp.MustCompile(`x(\d+)\/(\d+)`)
	rePartner  = regexp.MustCompile(`p(\w)\/(\w)`)
)

type Step struct {
	s string
}

type Steps []Step

func (s Step) Dance(str string) string {
	l := len(str)
	bs := []byte(str)
	if reSpin.MatchString(s.s) {
		matches := reSpin.FindStringSubmatch(s.s)
		val := atoi(matches[1])
		str = str[l-val:] + str[0:l-val]
	} else if reExchange.MatchString(s.s) {
		matches := reExchange.FindStringSubmatch(s.s)
		i1 := atoi(matches[1])
		i2 := atoi(matches[2])
		bs[i2], bs[i1] = bs[i1], bs[i2]
		str = string(bs)
	} else if rePartner.MatchString(s.s) {
		matches := rePartner.FindStringSubmatch(s.s)
		i1 := strings.Index(str, matches[1])
		i2 := strings.Index(str, matches[2])
		bs[i2], bs[i1] = bs[i1], bs[i2]
		str = string(bs)
	}
	return str
}

func (s Steps) Dance(str string) string {
	for _, step := range s {
		str = step.Dance(str)
	}
	return str
}

func main() {
	steps := getData("../data.txt")
	fmt.Println(steps.Dance(`abcdefghijklmnop`))
}

func getData(filename string) Steps {
	content, _ := file.GetContent(filename)
	tokens := strings.Split(strings.TrimSpace(string(content)), ",")
	steps := make(Steps, len(tokens))
	for i, token := range tokens {
		steps[i] = Step{s: token}
	}
	return steps
}

func atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}
