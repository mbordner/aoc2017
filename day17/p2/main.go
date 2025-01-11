package main

import "fmt"

func main() {
	p := 0
	val := 0

	// each time through, memory length will be i.  the first time through
	// the length will be 1 with just 0. to run it 50,000,000 times we just need to
	// run 1..50,000,000
	for i := 1; i <= 50000000; i++ {
		p = (p + 316) % i
		if p == 0 { // we only care about the value inserted after 0, which will always be index 1
			// as the 0 will always be at index 0, we don't care about any other value, so we won't keep the
			// memory
			val = i
		}
		p++
	}

	fmt.Println(val)
}
