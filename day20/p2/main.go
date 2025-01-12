package main

import (
	"fmt"
	"github.com/mbordner/aoc2017/common"
	"github.com/mbordner/aoc2017/common/file"
	"regexp"
	"strconv"
	"strings"
)

var (
	reParticle = regexp.MustCompile(`p=<(\s*-?\d+\s*),(\s*-?\d+\s*),(\s*-?\d+\s*)>, v=<(\s*-?\d+\s*),(\s*-?\d+\s*),(\s*-?\d+\s*)>, a=<(\s*-?\d+\s*),(\s*-?\d+\s*),(\s*-?\d+\s*)>`)
)

type V struct {
	X int
	Y int
	Z int
}

func (v V) Dis(o V) int {
	return abs(v.X-o.X) + abs(v.Y-o.Y) + abs(v.Z-o.Z)
}

func (v V) Add(o V) V {
	return V{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

type Particle struct {
	id  int
	pos V
	vel V
	acc V
}

type Particles []*Particle

func main() {
	//origin := V{}
	particles := getData("../data.txt")

	existing := make(map[int]*Particle)
	for _, p := range particles {
		existing[p.id] = p
	}

	changing := true
	for changing {

		prevPositions := make(map[int]V)
		positions := make(map[V]int)

		destroyed := make(map[int]bool)

		// advance positions
		for id, p := range existing {
			prevPositions[id] = p.pos // store position before change
			p.vel = p.vel.Add(p.acc)
			p.pos = p.pos.Add(p.vel)

			if pId, e := positions[p.pos]; e {
				destroyed[pId] = true
				destroyed[id] = true
			} else {
				positions[p.pos] = id
			}
		}

		// remove destroyed
		for id := range destroyed {
			delete(existing, id)
		}

		// now existing just has remaining particles that haven't collided, previously or in this round
		if len(existing) <= 1 {
			changing = false
		} else {
			existingIds := make([]int, 0, len(existing))
			for id := range existing {
				existingIds = append(existingIds, id)
			}

			anyDecreasingDistance := false
			pairs := common.GetPairSets(existingIds)
			for _, pair := range pairs {
				d1 := prevPositions[pair[0]].Dis(prevPositions[pair[1]])
				d2 := existing[pair[0]].pos.Dis(existing[pair[1]].pos)
				if d2 < d1 {
					anyDecreasingDistance = true
					break
				}
			}

			if !anyDecreasingDistance {
				changing = false
			}
		}
	}

	fmt.Println(len(existing))
}

func getData(filename string) Particles {
	lines, _ := file.GetLines(filename)
	particles := make(Particles, len(lines))

	for i, line := range lines {
		if reParticle.MatchString(line) {
			matches := reParticle.FindStringSubmatch(line)
			p := Particle{id: i}

			p.pos.X = atoi(matches[1])
			p.pos.Y = atoi(matches[2])
			p.pos.Z = atoi(matches[3])

			p.vel.X = atoi(matches[4])
			p.vel.Y = atoi(matches[5])
			p.vel.Z = atoi(matches[6])

			p.acc.X = atoi(matches[7])
			p.acc.Y = atoi(matches[8])
			p.acc.Z = atoi(matches[9])

			particles[i] = &p
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
	val, _ := strconv.Atoi(strings.TrimSpace(s))
	return val
}
