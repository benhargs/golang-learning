package main

import (
	"math"
	"math/rand"
)

func mathsEx() {
	pl("5 + 4=", 5+4)
	pl("5 - 4=", 5-4)
	pl("5 * 4=", 5*4)
	pl("5 / 4=", 5/4)
	pl("5 % 4=", 5%4)
	five := 5
	pl("five", five)
	five++
	pl("five++", five)
	five = 5
	five--
	pl("five--", five) //minus 1
	five = 5
	five += 3
	pl("five += 3:", five) // adds to the left what's on the right
	five = 5
	five -= 3
	pl("five -= 3:", five)

	randNum := rand.Intn(50) //generates a random number between 0 & 49
	pl(randNum)

	pl("Abs(-10) =", math.Abs(-10))   // Makes any number the positive of it.
	pl("Pow(4, 2) =", math.Pow(4, 2)) //4^2
	pl("Sqrt(16) =", math.Sqrt(16))
	pl("Cbrt(8) =", math.Cbrt(8))
	pl("Ceil(4.4) =", math.Ceil(4.4))   //Rounds up
	pl("Floor(4.4) =", math.Floor(4.4)) // Rounds down
	pl("Round(4.4) =", math.Round(4.4)) // Rounds normally
	pl("log2(8) =", math.Log2(8))       //log, calling the base in the function.
	pl("Log10(100) =", math.Log10(100))
	//Get the log of e to the power of 2
	pl("Log(7.389) =", math.Log(math.Exp(2)))
	numbArray := [3]int{15, -1, 6} //Declare length, then type, then fill in the data
	pl(numbArray)

	pl("Max(5,4) =", math.Max(5, 4)) //Finds the bigger value of two values
	pl("Min(5,4) =", math.Min(5, 4)) //Finds the smaller value of two

}
