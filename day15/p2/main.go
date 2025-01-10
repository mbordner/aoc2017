package main

import (
	"fmt"
	"sync"
)

type Generator struct {
	factor   uint32
	multiple uint32
	initial  uint32
	last     uint32
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

func NewGenerators(a, am, b, bm uint32) *Generators {
	return &Generators{mod: 2147483647, a: &Generator{factor: 16807, initial: a, last: a, multiple: am}, b: &Generator{factor: 48271, initial: b, last: b, multiple: bm}, productions: []Production{}, matches: []int{}}
}

func (g *Generator) Produce(wg *sync.WaitGroup, mod uint32, vals []uint32) {
	i := 0
	for i < len(vals) {
		val := g.Next(mod)
		if val%g.multiple == 0 {
			vals[i] = val
			i++
		}
	}
	wg.Done()
}

func (g *Generators) Produce() {
	pairsCount := 5000000
	as := make([]uint32, pairsCount)
	bs := make([]uint32, pairsCount)

	var wg sync.WaitGroup
	wg.Add(2)

	go g.a.Produce(&wg, g.mod, as)
	go g.b.Produce(&wg, g.mod, bs)

	wg.Wait()

	for i := 0; i < pairsCount; i++ {
		p := Production{a: as[i], b: bs[i]}
		if g.isMatch(p) {
			g.matches = append(g.matches, len(g.productions))
		}
		g.productions = append(g.productions, p)
	}

}

func (g *Generators) isMatch(p Production) bool {
	mask := uint32((1 << 16) - 1)
	a := p.a & mask
	b := p.b & mask
	return a == b
}

func main() {
	//gs := NewGenerators(65, 4, 8921, 8)
	gs := NewGenerators(634, 4, 301, 8)

	gs.Produce()

	fmt.Println(len(gs.matches))
}
