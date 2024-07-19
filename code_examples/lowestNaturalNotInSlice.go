package main

import (
	"fmt"
	"sort"
)

var pl = fmt.Println
var A = []int{1, 3, -7}

func Solution(A []int) int {
	sort.Ints(A)

	positiveCounter := 1

	for _, value := range A {
		if value < 1 {
			continue
		}

		if value == positiveCounter {
			positiveCounter++
			continue
		}

		if value > positiveCounter {
			return positiveCounter
		}
	}
	return positiveCounter
}

var lowest = Solution(A)

func main() {
	pl(lowest)
	offWork()
}
