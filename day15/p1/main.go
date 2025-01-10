package main

import "fmt"

type Generator struct {
	factor  uint32
	initial uint32
	last    uint32
}

func (g *Generator) Next(mod uint32) uint32 {
	g.last = uint32((uint64(g.last) * uint64(g.factor)) % uint64(mod))
	return g.last
}

type Production struct {
	a uint32
	b uint32
}

type Generators struct {
	a *Generator
	b *Generator

	mod uint32

	matches     []int
	productions []Production
}

func NewGenerators(a, b uint32) *Generators {
	return &Generators{mod: 2147483647, a: &Generator{factor: 16807, initial: a, last: a}, b: &Generator{factor: 48271, initial: b, last: b}, productions: []Production{}, matches: []int{}}
}

func (g *Generators) Produce() Production {
	p := Production{a: g.a.Next(g.mod), b: g.b.Next(g.mod)}
	if g.isMatch(p) {
		g.matches = append(g.matches, len(g.productions))
	}
	g.productions = append(g.productions, p)
	return p
}

func (g *Generators) isMatch(p Production) bool {
	mask := uint32((1 << 16) - 1)
	a := p.a & mask
	b := p.b & mask
	return a == b
}

func main() {
	gs := NewGenerators(634, 301)

	for i := 0; i < 40000000; i++ {
		gs.Produce()
	}

	fmt.Println(len(gs.matches))
}
