package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"regexp"
	"strconv"
)

var (
	reComponent = regexp.MustCompile(`(\d+)\/(\d+)`)
)

type Component []int
type ComponentList []Component

type ComponentsList []ComponentList

func (c Component) String() string {
	return fmt.Sprintf("[%d,%d]", c[0], c[1])
}

func (c Component) Strength() int {
	return c[0] + c[1]
}

func (cl ComponentList) Strength() int {
	strength := 0
	for _, c := range cl {
		strength += c.Strength()
	}
	return strength
}

func (cl ComponentList) Without(component Component) ComponentList {
	ocs := make(ComponentList, 0, len(cl)-1)
	for _, c := range cl {
		if !(c[0] == component[0] && c[1] == component[1]) {
			ocs = append(ocs, c)
		}
	}
	return ocs
}

func (cl ComponentList) Map() map[int]ComponentList {
	m := make(map[int]ComponentList)
	for _, c := range cl {
		if mcs, e := m[c[0]]; e {
			m[c[0]] = append(mcs, c)
		} else {
			m[c[0]] = ComponentList{c}
		}
		if c[1] != c[0] {
			if mcs, e := m[c[1]]; e {
				m[c[1]] = append(mcs, c)
			} else {
				m[c[1]] = ComponentList{c}
			}
		}
	}
	return m
}

func connect(head Component, orientation Component, avail ComponentList) ComponentsList {
	components := make(ComponentsList, 0, len(avail))

	am := avail.Map()
	if acs, e := am[orientation[1]]; e {
		for _, a := range acs {
			ao := a
			if a[0] != orientation[1] {
				ao = Component{a[1], a[0]}
			}
			tcs := connect(a, ao, avail.Without(a))
			for _, t := range tcs {
				components = append(components, append(ComponentList{head}, t...))
			}
		}
	}
	components = append(components, ComponentList{head})

	return components
}

func main() {
	avail := getComponents("../data.txt")
	availMap := avail.Map()

	var connectedComponents ComponentsList

	for _, r := range availMap[0] {
		connectedComponents = append(connectedComponents, connect(r, r, avail.Without(r))...)
	}

	maxStrength := 0

	for _, l := range connectedComponents {
		s := l.Strength()
		if s > maxStrength {
			maxStrength = s
		}
	}

	fmt.Println(maxStrength)
}

func getComponents(filename string) ComponentList {
	lines, _ := file.GetLines(filename)
	components := make(ComponentList, len(lines))
	for i, line := range lines {
		matches := reComponent.FindStringSubmatch(line)
		components[i] = Component{atoi(matches[1]), atoi(matches[2])}
	}
	return components
}

func atoi(s string) int {
	val, _ := strconv.ParseInt(s, 10, 64)
	return int(val)
}
