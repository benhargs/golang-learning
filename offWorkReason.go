package main

import (
	"fmt"
	"math/rand"
)

func offWork() {
	// Won't work on the Playground since the time is frozen.
	reasons := []string{
		"Locked out",
		"Pipes broke",
		"Food poisoning",
		"Not feeling well",
	}
	n := rand.Int() % len(reasons)
	fmt.Print("Gonna work from home...", reasons[n])
}
