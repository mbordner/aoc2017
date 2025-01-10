package main

import (
	"fmt"
	"strconv"
	"strings"
)

type KnotHash struct {
	data []byte
	skip int
	ptr  int
}

func NewKnotHash(hashStr string) *KnotHash {
	suffix := []int{17, 31, 73, 47, 23}
	input := make([]byte, 0, len(hashStr)+len(suffix))
	for i := range hashStr {
		input = append(input, byte(hashStr[i]))
	}
	for i := range suffix {
		input = append(input, byte(suffix[i]))
	}
	size := 256
	kh := &KnotHash{data: make([]byte, size)}
	for i := 0; i < size; i++ {
		kh.data[i] = byte(i)
	}
	for i := 0; i < 64; i++ {
		kh.process(input)
	}
	return kh
}

func (l *KnotHash) process(input []byte) {
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

func (l *KnotHash) advance(amt int) {
	ptr := l.ptr + amt
	for ptr >= len(l.data) {
		ptr -= len(l.data)
	}
	l.ptr = ptr
}

func (l *KnotHash) getData(length int) []byte {
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

func (l *KnotHash) setData(data []byte) {
	for i, j := l.ptr, 0; j < len(data); i, j = i+1, j+1 {
		if i >= len(l.data) {
			for i >= len(l.data) {
				i -= len(l.data)
			}
		}
		l.data[i] = data[j]
	}
}

func (l *KnotHash) reverse(data []byte) {
	for i, j, h := 0, len(data)-1, len(data)/2; i < h; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func (l *KnotHash) String() string {
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

func getBits(hash string) []bool {
	bits := make([]bool, 128)
	for b, h := 0, 0; h < len(hash); h++ {
		val, _ := strconv.ParseUint(hash[h:h+1], 16, 4)
		var bs []uint64
		for i := 0; i < 4; i++ {
			v := 1 & val
			bs = append([]uint64{v}, bs...)
			val >>= 1
		}
		for i := 0; i < 4; i, b = i+1, b+1 {
			if bs[i] == 1 {
				bits[b] = true
			} else {
				bits[b] = false
			}
		}
	}
	return bits
}

type Used struct {
	y int
	x int
}

type Regions map[Used]int

func (r Regions) Has(u Used) int {
	if v, e := r[u]; e {
		return v
	}
	return -1
}

func expandRegion(allUsed Regions, regions Regions, used Used, region int) {
	if _, isUsed := allUsed[used]; isUsed {
		if _, isExpanded := regions[used]; !isExpanded {
			regions[used] = region
			for _, d := range []Used{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
				expandRegion(allUsed, regions, Used{y: used.y + d.y, x: used.x + d.x}, region)
			}
		}
	}
}

func main() {
	hashInput := `ugkiagan`

	//bits := getBits(`a0c2017`)
	//fmt.Println(bits)
	//fmt.Println(NewKnotHash(``).String())
	//fmt.Println(NewKnotHash(`AoC 2017`).String())

	grid := make([][]byte, 128)
	for y := range grid {
		grid[y] = make([]byte, 128)
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}

	allUsed := make(Regions)
	regions := make(Regions)
	region := 0 // 0 is global for all used

	for y := 0; y < len(grid); y++ {
		kh := NewKnotHash(fmt.Sprintf("%s-%d", hashInput, y))
		for x, b := range getBits(kh.String()) {
			if b {
				grid[y][x] = '#'
				allUsed[Used{y, x}] = region
			}
		}
	}

	for u, v := range allUsed {
		if v >= 0 {
			if _, e := regions[u]; !e { // if we have not recorded this region
				region++
				expandRegion(allUsed, regions, u, region)
			}
		}
	}

	fmt.Printf("number of used sections: %d\n", len(allUsed))
	fmt.Printf("number of regions: %d\n", region)

}
