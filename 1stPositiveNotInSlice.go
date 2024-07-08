package main

import "sort"

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
