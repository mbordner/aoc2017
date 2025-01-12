package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common/file"
	"math"
	"regexp"
	"strconv"
)

var (
	reParticle = regexp.MustCompile(`p=<(-?\d+),(-?\d+),(-?\d+)>, v=<(-?\d+),(-?\d+),(-?\d+)>, a=<(-?\d+),(-?\d+),(-?\d+)>`)
)

type V struct {
	X int
	Y int
	Z int
}

func (v V) Dis(o V) int {
	return abs(v.X-o.X) + abs(v.Y-o.Y) + abs(v.Z-o.Z)
}

type Particle struct {
	pos V
	vel V
	acc V
}

type Particles []Particle

func main() {
	origin := V{}
	particles := getData("../data.txt")

	minDist := math.MaxUint32
	minPart := 0

	for i, p := range particles {
		d := p.acc.Dis(origin)
		if d < minDist {
			minDist = d
			minPart = i
		}
	}

	fmt.Println(minPart)
}

func getData(filename string) Particles {
	lines, _ := file.GetLines(filename)
	particles := make(Particles, len(lines))

	for i, line := range lines {
		if reParticle.MatchString(line) {
			matches := reParticle.FindStringSubmatch(line)
			p := Particle{}

			p.pos.X = atoi(matches[1])
			p.pos.Y = atoi(matches[2])
			p.pos.Z = atoi(matches[3])

			p.vel.X = atoi(matches[4])
			p.vel.Y = atoi(matches[5])
			p.vel.Z = atoi(matches[6])

			p.acc.X = atoi(matches[7])
			p.acc.Y = atoi(matches[8])
			p.acc.Z = atoi(matches[9])

			particles[i] = p
		} else {
			panic("invalid line")
		}
	}
	return particles
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}
