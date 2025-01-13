package main

import (
	"fmt"
	"math"
)

func prime(n int64) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}
	sqrt := int64(math.Sqrt(float64(n)))
	for i := int64(3); i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 914 too low
func main() {

	b := int64(105700)
	c := b + 17000

	h := int64(0)

	for b <= c {
		if !prime(b) {
			h++
		}

		b += 17
	}

	fmt.Println(h)

}
